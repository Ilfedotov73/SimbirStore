import { formatPrice } from "@/utils/format";
import { getIconByRating } from "@/utils/rating";
import { Link } from "react-router";

interface ProductCardProps {
    id: number,
    name: string,
    price: number,
    imageUrl: string,
    rating: number
}

export default function ProductCard({id, name, price, imageUrl, rating} : ProductCardProps) {
    return (
        <Link to={`/products/${id}`} className="space-y-1 w-50">
            <img src={imageUrl} alt="product" className="w-50 h-55 rounded-2xl"/>
            <div className="flex justify-between space-x-1">
                <p className="text-wrap">{name}</p>
                <div className="flex mb-auto items-center space-x-0.5">
                    {getIconByRating(rating)}
                    <p>{rating}</p>
                </div>
                
            </div>
            
            <p className="font-semibold text-xl">{formatPrice(price)}</p>
        </Link>
    );
}