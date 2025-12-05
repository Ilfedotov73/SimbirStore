import ReviewCard from "@/components/pages/review-card";
import { Pagination, PaginationContent, PaginationEllipsis, PaginationItem, PaginationLink } from "@/components/ui/pagination";
import { ChevronLeftIcon } from "lucide-react";
import { Link, useParams } from "react-router";

function Reviews() {
    const params = useParams();
    const test = Array.from({ length: 12 }, (_, index) => index + 1);

    return (
        <div className="mx-auto space-y-4">
             <div className="space-y-2">
                <Link to={`/products/${params.productId}`} className="flex space-x-4 items-center">
                    <h1 className="font-semibold text-2xl">Отзывы на Телефон iPhone 15 Pro Max</h1>
                    <ChevronLeftIcon />
                </Link>
                <p>Всего: 124</p>
            </div>
            <Pagination>
                <PaginationContent>
                    <PaginationItem>
                        <PaginationLink href="#" isActive>1</PaginationLink>
                    </PaginationItem>
                    <PaginationItem>
                        <PaginationLink href="#" >2</PaginationLink>
                    </PaginationItem>
                    <PaginationItem>
                        <PaginationLink href="#">3</PaginationLink>
                    </PaginationItem>
                    <PaginationItem>
                        <PaginationEllipsis />
                    </PaginationItem>
                    <PaginationItem>
                        <PaginationLink href="#">7</PaginationLink>
                    </PaginationItem>
                    <PaginationItem>
                        <PaginationLink href="#">8</PaginationLink>
                    </PaginationItem>
                    <PaginationItem>
                        <PaginationLink href="#">9</PaginationLink>
                    </PaginationItem>
                </PaginationContent>
            </Pagination>
            <div className="space-y-6">
                {
                    test.map((item) => (
                        <ReviewCard key={item} authorFullName={"Иванов Иван"} authorAvatarUrl={"https://placehold.co/200x200"} text={"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."} rating={4}/>
                    ))
                }
            </div>
        </div>
    );
}

export default Reviews;