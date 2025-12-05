import NotificationCard from "@/components/pages/notification-card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Pagination, PaginationContent, PaginationEllipsis, PaginationItem, PaginationLink } from "@/components/ui/pagination";

function Notifications() {
    const test = Array.from({ length: 12 }, (_, index) => index + 1);

    return (
        <div className="mx-auto space-y-8">
            <div className="space-y-2 self-start">
                <h2 className="font-semibold text-2xl">Уведомления</h2>
                <p>Всего: 124</p>
            </div>
            <div>
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
            </div>
            <div className="flex">
                <div className="space-y-4 w-50">
                    <div className="space-y-2">
                        <Label htmlFor="start">Начиная</Label>
                        <Input name="price-from" type="date" placeholder="с"/>
                    </div>
                    <div className="space-y-2">
                        <Label htmlFor="end">Заканчивая</Label>
                        <Input name="price-from" type="date" placeholder="до"/>
                    </div>
                    <Button className="w-full">Применить</Button>
                </div>
                <div className="space-y-6 ml-16">
                    {test.map((item) => (
                        <NotificationCard key={item} text="Текст уведомления" date={new Date()}/>
                    ))}
                </div>
            </div>
        </div>
    );
}

export default Notifications;