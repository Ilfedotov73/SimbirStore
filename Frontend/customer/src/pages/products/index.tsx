import ProductCard from "@/components/pages/product-card";
import { Button } from "@/components/ui/button";
import { Command, CommandGroup, CommandItem } from "@/components/ui/command";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Pagination, PaginationContent, PaginationEllipsis, PaginationItem, PaginationLink } from "@/components/ui/pagination";
import { Popover, PopoverContent, PopoverTrigger } from "@/components/ui/popover";
import { cn } from "@/lib/utils";
import { Check, ChevronsUpDown } from "lucide-react";
import { useState } from "react";
import { useSearchParams } from "react-router";

function Products() {
    const [params, setParams] = useSearchParams();

    const [openSort, setOpenSort] = useState(false);
    const [openVendor, setOpenVendor] = useState(false);
    const [valueSort, setValueSort] = useState("price-asc");
    const [valueVendor, setValueVendor] = useState("");

    const test = Array.from({ length: 12 }, (_, index) => index + 1);

    return (
        <div className="mx-auto space-y-8">
            <div className="space-y-2 self-start">
                <h2 className="font-semibold text-2xl">Товары по запросу: {params.get("q")}</h2>
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
                        <Label htmlFor="sort">Сортировать по</Label>
                        <Popover open={openSort} onOpenChange={setOpenSort}>
                            <PopoverTrigger asChild>
                                <Button
                                    id="sort"
                                    variant="outline"
                                    role="combobox"
                                    aria-expanded={openSort}
                                    className="justify-between w-50"
                                    >
                                    {valueSort}
                                    <ChevronsUpDown className="opacity-50" />
                                </Button>
                            </PopoverTrigger>
                            <PopoverContent className="p-0 w-50">
                                <Command>
                                    <CommandGroup>
                                        <CommandItem
                                            key="sort1"
                                            value="price-asc"
                                            onSelect={(c) => {
                                                setValueSort(c === valueSort ? "" : c)
                                                setOpenSort(false)
                                            }}>
                                            Сначала дешевые
                                            <Check
                                                className={cn(
                                                "ml-auto",
                                                valueSort === "price-asc" ? "opacity-100" : "opacity-0"
                                                )}/>
                                        </CommandItem>
                                        <CommandItem
                                            key="sort2"
                                            value="price-desc"
                                            onSelect={(c) => {
                                                setValueSort(c === valueSort ? "" : c)
                                                setOpenSort(false)
                                            }}>
                                            Сначала дорогие
                                            <Check
                                                className={cn(
                                                "ml-auto",
                                                valueSort === "price-desc" ? "opacity-100" : "opacity-0"
                                                )}/>
                                        </CommandItem>
                                    </CommandGroup>
                                </Command>
                            </PopoverContent>
                        </Popover>
                    </div>
                    <div className="space-y-2">
                        <Label htmlFor="price">Цена</Label>
                        <div id="price" className="flex space-x-2">
                            <Input name="price-from" type="number" min="0" placeholder="от"/>
                            <Input name="price-to" type="number" min="0" placeholder="до"/>
                        </div>
                    </div>
                    <div className="space-y-2">
                        <Label htmlFor="vendor">Продавец</Label>
                        <Popover open={openVendor} onOpenChange={setOpenVendor}>
                            <PopoverTrigger asChild>
                                <Button
                                    id="vendor"
                                    variant="outline"
                                    role="combobox"
                                    aria-expanded={openVendor}
                                    className="justify-between w-full"
                                    >
                                    {valueVendor}
                                    <ChevronsUpDown className="opacity-50" />
                                </Button>
                            </PopoverTrigger>
                            <PopoverContent className="p-0 w-50">
                                <Command>
                                    <CommandGroup>
                                        <CommandItem
                                            key="vendor1"
                                            value="Test 1"
                                            onSelect={(c) => {
                                                setValueVendor(c === valueVendor ? "" : c)
                                                setOpenSort(false)
                                            }}>
                                            Test 1
                                            <Check
                                                className={cn(
                                                "ml-auto",
                                                valueVendor === "test 1" ? "opacity-100" : "opacity-0"
                                                )}/>
                                        </CommandItem>
                                        <CommandItem
                                            key="vendor2"
                                            value="Test 2"
                                            onSelect={(c) => {
                                                setValueVendor(c === valueVendor ? "" : c)
                                                setOpenSort(false)
                                            }}>
                                            Test 2
                                            <Check
                                                className={cn(
                                                "ml-auto",
                                                valueVendor === "test 2" ? "opacity-100" : "opacity-0"
                                                )}/>
                                        </CommandItem>
                                    </CommandGroup>
                                </Command>
                            </PopoverContent>
                        </Popover>
                    </div>
                    <Button className="w-full">Применить</Button>
                </div>
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

export default Products;