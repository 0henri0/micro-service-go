package services

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"log"
	"micro-service-go/internal/models"
	"strconv"
	"strings"
	"sync"
	"time"
)

type ScrappingHtml interface {
	GetOuterHTML(url string, excludeUrls []string) (string, error)
	GetUrlList(outerHtml string) ([]string, error)
	GetTotalPage(outerHtml string) int
	GetProductInfoFromList(outerHtml string) ([]models.Product, error)
}

type scrappingHtml struct {
}

func NewScrappingHtml() ScrappingHtml {
	return &scrappingHtml{}
}

func chromeDpContext() (context.Context, context.CancelFunc) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("start-fullscreen", false),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-extensions", false),
	)
	ctxWithTimeout, _ := context.WithTimeout(context.Background(), time.Second*50)
	initialCtx, _ := chromedp.NewExecAllocator(ctxWithTimeout, opts...)
	ctx, cancel := chromedp.NewContext(initialCtx, chromedp.WithLogf(log.Printf))

	return ctx, cancel
}

func (s *scrappingHtml) GetOuterHTML(url string, excludeUrls []string) (string, error) {
	var outerHtml string
	ctx, cancel := chromeDpContext()
	defer cancel()
	fmt.Println(excludeUrls)
	task := chromedp.Tasks{
		network.SetBlockedURLS(excludeUrls),
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(ctx context.Context) error {
			node, err := dom.GetDocument().Do(ctx)

			if err != nil {
				return err
			}

			res, err := dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			if err != nil {
				return err
			}
			outerHtml = res
			return nil
		})}

	if err := chromedp.Run(ctx, task); err != nil {
		fmt.Println(err)

		return "", err
	}

	return outerHtml, nil
}

func (s *scrappingHtml) GetTotalPage(outerHtml string) int {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(outerHtml))
	if err != nil {
		return 0
	}
	last := doc.Find(".ant-pagination-item").Last()
	totalPage, exist := last.Attr("title")
	if !exist {
		return 0
	}
	totalPageInt, err := strconv.Atoi(totalPage)
	if err != nil {
		return 0
	}

	return totalPageInt
}

func (s *scrappingHtml) GetUrlList(outerHtml string) ([]string, error) {
	url := make([]string, 0)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(outerHtml))
	if err != nil {
		return nil, err
	}

	doc.Find("div[data-qa-locator='product-item']").Each(func(index int, query *goquery.Selection) {
		text, _ := query.Find("a").First().Attr("href")
		url = append(url, text)
	})

	return url, nil
}

func (s *scrappingHtml) GetProductInfoFromList(outerHtml string) ([]models.Product, error) {
	var wg sync.WaitGroup
	products := make([]models.Product, 0)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(outerHtml))
	if err != nil {
		return nil, err
	}
	numberProducts := 0
	var queries []*goquery.Selection
	doc.Find("div[data-qa-locator='product-item']").Each(func(index int, query *goquery.Selection) {
		numberProducts++
		queries = append(queries, query)
	})
	productsChannel := make(chan models.Product, numberProducts)

	for _, query := range queries {
		wg.Add(1)
		go getProductInfo(query, productsChannel, &wg)
	}

	wg.Wait()

	// Close the channel to signal to the receiving goroutine that we're done
	close(productsChannel)

	for i := 0; i < numberProducts; i++ {
		product := <-productsChannel
		products = append(products, product)
	}

	return products, nil
}

func getProductInfo(query *goquery.Selection, productChannel chan<- models.Product, wg *sync.WaitGroup) {
	defer wg.Done()
	selection := query.Find("a").Last()
	originUrl, err := selection.Attr("href")
	if !err {
		originUrl = ""
	} else {
		originUrl = strings.Trim(originUrl, " ")
		originUrl = strings.Trim(originUrl, "/")
		originUrl = strings.Trim(originUrl, "www.")
		originUrl = strings.Trim(originUrl, "https://")
		originUrl = strings.Trim(originUrl, "http://")
	}
	text, _ := selection.Attr("title")
	productChannel <- models.Product{
		OriginUrl: originUrl,
		Name:      text,
	}
}
