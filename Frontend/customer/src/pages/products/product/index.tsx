import { Button } from "@/components/ui/button";
import { parseDesriptionToElements } from "@/utils/descr";
import { formatPrice } from "@/utils/format";
import { getIconByRating } from "@/utils/rating";
import { CalendarIcon, ChevronRight, PhoneIcon, UserIcon } from "lucide-react";
import { Link, useParams } from "react-router";

function Product() {
    const params = useParams();
    return (
        <div className="mx-auto flex space-x-20">
            <div className="space-y-6">
                <div className="space-y-2">
                    <h1 className="font-semibold text-2xl">Телефон iPhone 15 Pro Max</h1>
                    <div className="flex space-x-4">
                        <div className="flex space-x-1 items-center">
                            {getIconByRating(4.7)}
                            <p>{4.7}</p>
                        </div>
                        <Link to={`/products/${params.productId}/reviews`} className="text-gray-500 flex space-x-2  items-center">
                            <p>134 оценок</p>
                            <ChevronRight size={16}/>
                        </Link>
                    </div>
                    <div className="mt-4">
                        <img src="https://placehold.co/500x300" className="rounded-2xl"/>
                    </div>
                </div>
                <div className="space-y-4">
                    <h2 className="font-semibold text-2xl">Характеристики</h2>
                    <div className="grid grid-cols-2 gap-x-16 gap-y-4">
                        {parseDesriptionToElements("Цвет:черный;Память:128 GB;Производитель:Apple;Модель:iPhone 15 Pro Max")}
                    </div>
                </div>
                <div className="space-y-4">
                    <h2 className="font-semibold text-2xl">Об объявлении</h2>
                    <div className="grid grid-cols-2 gap-x-16 gap-y-4">
                        <div className="flex justify-between col-1">
                            <span className="text-gray-500">ID</span>
                            <span>1</span>
                        </div>
                        <div className="flex justify-between col-1">
                            <span className="text-gray-500">Выложено</span>
                            <span>{new Date().toLocaleDateString()}</span>
                        </div>
                    </div>
                </div>
            </div>
            <div className="space-y-5 w-50">
                <div className="space-y-3">
                    <p className="font-semibold text-2xl">{formatPrice(65000)}</p>
                    <Button>Предложить сделку</Button>
                </div>
                <div className="space-y-4">
                    <h2 className="font-semibold text-2xl">Продавец</h2>
                    <div className="space-y-1.5">
                        <div className="flex space-x-2 items-center">
                            <UserIcon size={16} className="text-gray-500"/>
                            <p>Иванов Иван</p>
                        </div>
                        <div className="flex space-x-2 items-center">
                            <PhoneIcon size={16} className="text-gray-500"/>
                            <p>+7 999 123 45 67</p>
                        </div>
                        <div className="flex space-x-2 items-center">
                            <PhoneIcon size={16} className="text-gray-500"/>
                            <p>@vendor_ivan</p>
                        </div>
                        <div className="flex space-x-2 items-center">
                            <CalendarIcon size={16} className="text-gray-500"/>
                            <p>1 день</p>
                        </div>
                    </div>
                    <Button asChild className="w-full">
                        <Link to={`/vendor/${1}`}>Профиль</Link>
                    </Button>
                </div>
            </div>
        </div>
    );
}

export default Product;