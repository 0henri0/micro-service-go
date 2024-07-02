'use client';

import React from 'react';

import { Navigation, Pagination, EffectFade } from 'swiper/modules';
import { Swiper, SwiperSlide } from 'swiper/react';

// Import Swiper styles
import 'swiper/css';
import 'swiper/css/navigation';
import 'swiper/css/pagination';
import 'swiper/css/effect-fade';

const Carousel: React.FC = () => {
  return (
    <div className="w-full mt-8">
      <Swiper
        autoHeight={true}
        spaceBetween={30}
        modules={[Navigation, Pagination, EffectFade]}
        effect={'fade'}
        slidesPerView={1}
        loop={true}
        navigation
        pagination={{ clickable: true }}
      >
        <SwiperSlide>
          <img src="https://cf.shopee.vn/file/70d870c35efa10f6e0acccd0db96df22" alt="Slide 1" className="w-full" />
        </SwiperSlide>
        <SwiperSlide>
          <img src="https://cf.shopee.vn/file/70d870c35efa10f6e0acccd0db96df22" alt="Slide 2" className="w-full" />
        </SwiperSlide>
        <SwiperSlide>
          <img src="https://cf.shopee.vn/file/70d870c35efa10f6e0acccd0db96df22" className="w-full" />
        </SwiperSlide>
        {/* Add more slides as needed */}
      </Swiper>
    </div>
  );
};

export default Carousel;
