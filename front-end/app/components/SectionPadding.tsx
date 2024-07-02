import React from "react";

interface SectionPadding {
    img: string;
    title: string;
    bgColor: string;
}

const sectionPaddings: SectionPadding[] = [
    {
        img: 'https://nathstore.site/frontend/assets/imgs/theme/icons/feature-1.png',
        title: 'Free Shipping',
        bgColor: 'bg-[#fddde4]'

    },
    {
        img: 'https://nathstore.site/frontend/assets/imgs/theme/icons/feature-2.png',
        title: 'Online Order',
        bgColor: 'bg-[#d1e8f2]'

    },
    {
        img: 'https://nathstore.site/frontend/assets/imgs/theme/icons/feature-3.png',
        title: 'Save Money',
        bgColor: 'bg-[#cdebbc]'

    },
    {
        img: 'https://nathstore.site/frontend/assets/imgs/theme/icons/feature-4.png',
        title: 'Promotions',
        bgColor: 'bg-[#cdd4f8]'

    },
    {
        img: 'https://nathstore.site/frontend/assets/imgs/theme/icons/feature-5.png',
        title: 'Happy Sell',
        bgColor: 'bg-[#f6dbf6]'

    },{
        img: 'https://nathstore.site/frontend/assets/imgs/theme/icons/feature-6.png',
        title: '24/7 Support',
        bgColor: 'bg-[#fff2e5]'

    }
]
const SectionPadding: React.FC = () => {
    return (
        <div className="grid grid-cols-2 gap-4 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-6 xl:gap-8 mt-4">
            {sectionPaddings.map((sectionPadding, idx) => (
                <div key={idx} className="h-60 rounded border-solid border border-[#cce7d0] pl-4 pr-4 pt-6 pb-6 flex flex-col justify-between"> 
                    <img src={sectionPadding.img} alt={sectionPadding.title} className="h-28 mx-auto" />
                    <div className="text-center">
                        <h3 className={`text-lg py-1 px-3 rounded inline-block ${sectionPadding.bgColor}`}>{sectionPadding.title}</h3>
                    </div>
                </div>
            ))}
        </div>
    )
};

export default SectionPadding;