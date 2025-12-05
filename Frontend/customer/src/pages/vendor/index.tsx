import ProductCard from "@/components/pages/product-card";
import { Pagination, PaginationContent, PaginationEllipsis, PaginationItem, PaginationLink } from "@/components/ui/pagination";
import { BoxIcon, CalendarIcon, PhoneIcon    } from "lucide-react";

function Vendor() {
    const test = Array.from({ length: 12 }, (_, index) => index + 1);

    return (
        <div className="mx-auto flex space-x-8">
            <div className="space-y-2 self-start">
                <img src="https://placehold.co/125x125" alt="avatar" className="rounded-2xl"/>
                <h2 className="font-semibold text-2xl">Иванов Иван</h2>
                <div className="space-y-1.5">
                    <div className="flex space-x-2 items-center">
                        <BoxIcon size={16} className="text-gray-500"/>
                        <p>124</p>
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
            </div>
            <div className="space-y-8">
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
                <div className="ml-10 mr-auto grid md:grid-cols-2 xl:grid-cols-4 gap-6">
                {
                    test.map((item) => <ProductCard key={item} 
                        id={item} 
                        name={"Телефон iPhone 15 Pro Max"} 
                        price={65000} 
                        imageUrl={"https://placehold.co/200x220"} 
                        rating={4.7} />)
                }
            </div>
            </div>
        </div>
    );
}

export default Vendor;