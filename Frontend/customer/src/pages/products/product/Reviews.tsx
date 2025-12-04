import { useParams } from "react-router";

function Reviews() {
    const params = useParams();
    return (
        <div>Reviews for: {params.productId}</div>
    );
}

export default Reviews;