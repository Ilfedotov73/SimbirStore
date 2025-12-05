import { Button } from "@/components/ui/button";
import { ButtonGroup } from "@/components/ui/button-group";
import { Input } from "@/components/ui/input";
import { BellIcon, SearchIcon, UserIcon } from "lucide-react";
import { Form, Link, Outlet } from "react-router";

export default function Layout() {
    return (
        <>
            <header className="flex mx-5 sm:mx-20 md:mx-30 xl:mx-60 h-20 items-center">
                <div className="mr-auto">
                    <Link to="/" className="font-semibold text-2xl">SimbirStore</Link>
                </div>
                <Form className="w-full max-w-40 sm:max-w-60 md:max-w-75 xl:max-w-130" action="/products" method="get">
                    <ButtonGroup className="w-full">
                        <Input type="search" name="q" placeholder="Поиск товаров" />
                        <Button variant="outline" size="icon">
                            <SearchIcon/>
                        </Button>
                    </ButtonGroup>
                </Form>
                <div className="ml-auto space-x-2 sm:space-x-4 md:space-x-9 xl:space-x-18">
                    <Button asChild variant="ghost" size="icon">
                        <Link to="/notifications">
                            <BellIcon/>
                        </Link>
                    </Button>
                    <Button asChild variant="ghost" size="icon">
                        <Link to="/profile">
                            <UserIcon/>
                        </Link>
                    </Button>
                </div>
            </header>
            <main className="min-h-screen flex flex-col items-center mt-12">
                <Outlet />
            </main>
        </>   
    )
}