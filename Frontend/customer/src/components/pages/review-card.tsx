import { getRatingIcons } from "@/utils/rating";

interface ReviewCardProps {
    authorFullName: string,
    authorAvatarUrl: string,
    text: string,
    rating: number
}

export default function ReviewCard({authorAvatarUrl, authorFullName, rating, text} : ReviewCardProps) {
    return (
        <div className="space-y-1 w-200">
            <div className="flex space-x-3">
                <img src={authorAvatarUrl} alt="avatar" className="w-12 h-12 rounded-xl"/>
                <div className="space-y-0.5">
                    <p>{authorFullName}</p>
                    {getRatingIcons(rating)}
                </div>
            </div>
            <p>{text}</p>   
        </div>
    );
}