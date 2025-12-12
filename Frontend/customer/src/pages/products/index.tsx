import type { ProductsDTOType } from "@/api/dtos/products";
import { getProducts } from "@/api/service/products";
import { PaginationComponent } from "@/components/helpers/pagination";
import ProductCard from "@/components/pages/product-card";
import { Button } from "@/components/ui/button";
import { Command, CommandGroup, CommandItem } from "@/components/ui/command";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Popover, PopoverContent, PopoverTrigger } from "@/components/ui/popover";
import { cn } from "@/lib/utils";
import { Check, ChevronsUpDown } from "lucide-react";
import { useEffect, useState } from "react";
import { useSearchParams } from "react-router";

const filters = [
    {
        value: "price.asc",
        label: "Сначала дешевые",
    },
    {
        value: "price.desc",
        label: "Сначала дорогие",
    },
    {
        value: "rating.asc",
        label: "Сначала с низкой оценкой",
    },
    {
        value: "rating.desc",
        label: "Сначала с высокой оценкой",
    },
];

function Products() {
    const [params, setParams] = useSearchParams();

    const [openSort, setOpenSort] = useState(false);
    const [openVendor, setOpenVendor] = useState(false);

    const [kvSort, setKvSort] = useState({value: "", label: ""});

    const [valueVendor, setValueVendor] = useState<string | undefined>(undefined);
    const [priceFrom, setPriceFrom] = useState<number | undefined>(undefined);
    const [priceTo, setPriceTo] = useState<number | undefined>(undefined);

    const [data, setData] = useState<ProductsDTOType | Error>();

    const q = params.get("q") ?? undefined;
    const page = Number(params.get("page") ?? 1);

    const handlePageChange = (page : number) => {
        params.set("page", page.toString());
        setParams(params);
    };

    const handleFilter = () => {
        console.log(kvSort, valueVendor, priceFrom, priceTo);

        const data = async () => {
            const res = await getProducts(page, priceFrom, priceTo, undefined, q, {
                field: kvSort.value.split(".")[0],
                order: kvSort.value.split(".")[1],
            });
            setData(res);
        }

        data();
    };

    useEffect(() => {
        const data = async () => {
            const res = await getProducts(page, undefined, undefined, undefined, q);
            setData(res);
        }

        data();
    }, [page,  q]);

    return (
        <div className="mx-auto w-[80%] space-y-8">
            <div className="space-y-2 self-start">
                <h2 className="font-semibold text-2xl">
                    {
                       params.get("q") == "" || params.get("q") == undefined ? "Все товары" : `Товары по запросу: ${params.get("q")}`
                    }</h2>
                <p>Всего: {data && !(data instanceof Error) ? data.paging.total : 0}</p>
            </div>
            <div>
                { data && !(data instanceof Error) && <PaginationComponent 
                    totalItems={data.paging.total}
                    currentPage={data.paging.page}
                    pageSize={data.paging.size}
                    handleOnPageClick={handlePageChange}
                    />
                }
            </div>
            <div className="flex">
                <div className="space-y-4 w-70">
                    <div className="space-y-2">
                        <Label htmlFor="sort">Сортировать по</Label>
                        <Popover open={openSort} onOpenChange={setOpenSort}>
                            <PopoverTrigger asChild>
                                <Button
                                    id="sort"
                                    variant="outline"
                                    role="combobox"
                                    aria-expanded={openSort}
                                    className="justify-between w-70"
                                    >
                                    {kvSort.label}
                                    <ChevronsUpDown className="opacity-50" />
                                </Button>
                            </PopoverTrigger>
                            <PopoverContent className="p-0 w-70">
                                <Command>
                                    <CommandGroup>
                                        {
                                            filters.map((filter) => (
                                                <CommandItem
                                                    key={filter.value}
                                                    value={filter.label}
                                                    onSelect={(c) => {
                                                        const selectedFilter = filters.find(f => f.label.toLowerCase() === c.toLowerCase());
                                                        if (filter.label === kvSort.label) {
                                                            setKvSort({value: "", label: ""}); 
                                                        } else {
                                                            setKvSort(selectedFilter || {value: "", label: ""});
                                                        }
                                                        
                                                        setOpenSort(false);
                                                    }}>
                                                    {filter.label}
                                                    <Check
                                                        className={cn(
                                                        "ml-auto",
                                                        kvSort.label === filter.label ? "opacity-100" : "opacity-0"
                                                        )}/>
                                                </CommandItem>
                                            ))
                                        }
                                    </CommandGroup>
                                </Command>
                            </PopoverContent>
                        </Popover>
                    </div>
                    <div className="space-y-2">
                        <Label htmlFor="price">Цена</Label>
                        <div id="price" className="flex space-x-2">
                            <Input name="price-from" type="number" min="0" placeholder="от" value={priceFrom} onChange={(e) => {return setPriceFrom(Number(e.target.value))}}/>
                            <Input name="price-to" type="number" min="0" placeholder="до" value={priceTo} onChange={(e) => {return setPriceTo(Number(e.target.value))}}/>
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
                    <Button className="w-full" onClick={handleFilter}>Применить</Button>
                </div>
                <div className="ml-10 mr-auto grid md:grid-cols-2 xl:grid-cols-4 gap-6">
                    {
                        data && (data instanceof Error) ? 
                        <p>{data.message}</p>
                        :
                        data && data.items.map((item) => <ProductCard key={item.id} 
                            id={item.id} 
                            name={item.name} 
                            price={item.price} 
                            imageUrl={item.photoUrl} 
                            rating={item.productRating} />)
                    }
                </div>
            </div>
        </div>
    );
}

export default Products;