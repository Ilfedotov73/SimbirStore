import { MdOutlineStar, MdOutlineStarBorder, MdOutlineStarHalf } from "react-icons/md";

export function getIconByRating(rating: number) {
    if (rating <= 2) {
        return (
            <MdOutlineStarBorder size={18}/>
        );
    }
    else if (rating < 5) {
        return (
            <MdOutlineStarHalf size={18}/>
        )
    } else {
        return (
            <MdOutlineStar size={18}/>
        )
    }
}

export function getRatingIcons(rating: number) {
    const icons = [];
    for (let i = 0; i < rating; i++) {
        icons.push(<MdOutlineStar size={18}/>);
    }
    for (let i = 0; i < 5 - rating; i++) {
        icons.push(<MdOutlineStarBorder size={18}/>)
    }
    return (
        <div className="flex space-x-0.5">
            {icons}
        </div>
    );
}