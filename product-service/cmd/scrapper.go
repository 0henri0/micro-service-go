package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := newChromedp()
	defer cancel()

	getUrlList("https://www.lazada.vn/catalog?q=quan", ctx)
}

func getUrlList(url string, ctx context.Context) {
	task := chromedp.Tasks{
		network.SetBlockedURLS([]string{
			"https://tpsservice-files-inner.cn-hangzhou.oss-cdn.aliyun-inc.com/images/resources/9dd6917e501f4144dd7af71009cceb63-1-1.png",
		}),
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

			doc, err := goquery.NewDocumentFromReader(strings.NewReader(res))
			if err != nil {
				return err
			}

			doc.Find("div[data-qa-locator='product-item']").Each(func(index int, query *goquery.Selection) {
				text, _ := query.Find("a").First().Attr("href")
				fmt.Println(text)
			})

			return nil
		}),
	}
	if err := chromedp.Run(ctx, task); err != nil {
		fmt.Println(err)
	}
}

func newChromedp() (context.Context, context.CancelFunc) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("start-fullscreen", false),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-extensions", false),
	)
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	allocCtx2, can := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))

	return allocCtx2, can
}
