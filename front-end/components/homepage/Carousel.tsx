"use client";

import React from "react";

import { Navigation, Pagination, EffectFade } from "swiper/modules";
import { Swiper, SwiperSlide } from "swiper/react";

// Import Swiper styles
import "swiper/css";
import "swiper/css/navigation";
import "swiper/css/pagination";
import "swiper/css/effect-fade";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  faChevronLeft,
  faChevronRight,
} from "@fortawesome/free-solid-svg-icons";

const Carousel: React.FC = () => {
  return (
    <div className="w-full mt-8">
      <Swiper
        autoHeight={true}
        spaceBetween={30}
        modules={[Navigation, Pagination, EffectFade]}
        effect={"fade"}
        slidesPerView={1}
        loop={true}
        navigation={{
          prevEl: ".prev",
          nextEl: ".next",
        }}
        pagination={{ clickable: true }}
        className="mySwiper"
      >
        <SwiperSlide>
          <img
            src="https://cf.shopee.vn/file/70d870c35efa10f6e0acccd0db96df22"
            alt="Slide 1"
            className="w-full"
          />
        </SwiperSlide>
        <SwiperSlide>
          <img
            src="https://cf.shopee.vn/file/70d870c35efa10f6e0acccd0db96df22"
            alt="Slide 2"
            className="w-full"
          />
        </SwiperSlide>
        <SwiperSlide>
          <img
            src="https://cf.shopee.vn/file/70d870c35efa10f6e0acccd0db96df22"
            className="w-full"
          />
        </SwiperSlide>
        {/* Add more slides as needed */}
        <div className="slider-arrow absolute top-[45%] z-10 left-0 w-full flex justify-between text-mint">
          <button aria-label="prev" type="button" className=" ml-5 prev cursor-pointer bg-mintLight h-11 w-11 rounded-full flex items-center justify-center hover:bg-mint hover:text-white">
            <FontAwesomeIcon icon={faChevronLeft} />
          </button>
          <button aria-label="next" type="button" className="mr-5 next cursor-pointer bg-mintLight h-11 w-11 rounded-full flex items-center justify-center hover:bg-mint hover:text-white">
            <FontAwesomeIcon icon={faChevronRight} />
          </button>
        </div>
      </Swiper>
    </div>
  );
};

export default Carousel;
