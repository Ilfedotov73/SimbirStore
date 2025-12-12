/* eslint-disable react-hooks/set-state-in-effect */
import { useEffect, useState } from "react";
import { Pagination, PaginationContent, PaginationEllipsis, PaginationItem } from "../ui/pagination";
import { range } from "lodash";
import { Button } from "../ui/button";

interface PaginationComponentProps {
    totalItems: number
    currentPage: number
    pageSize: number
    handleOnPageClick: (page: number) => void;
}

export function PaginationComponent({ totalItems, currentPage, pageSize, handleOnPageClick } : PaginationComponentProps) {
    const toShowButtons = 5;
    const [showStartEllipsis, setShowStartEllipsis] = useState(false);
    const [showEndEllipsis, setShowEndEllipsis] = useState(false);
    const [currButtons, setCurrButtons] = useState<number[]>([]);

    const totalPages = Math.ceil(totalItems / pageSize);

    useEffect(() => {
        if (currentPage === 1)
            setShowStartEllipsis(false);
    if (totalPages <= toShowButtons) {
        setCurrButtons(range(1, totalPages + 1));
    }
    else if (currentPage + toShowButtons - 1 < totalPages) {
        if (currentPage !== 1)
            setShowStartEllipsis(true);
        setShowEndEllipsis(true);
        setCurrButtons(range(currentPage, toShowButtons + currentPage));
    }
    else {
        setCurrButtons(range(totalPages - toShowButtons, totalPages + 1));
        setShowEndEllipsis(false);
    }
    }, [currentPage, totalPages])
    

    return (
        <Pagination>
            <PaginationContent>
                <PaginationItem>
                    <Button variant="ghost" onClick={() => handleOnPageClick(currentPage - 1)} disabled={currentPage <= 1}>Назад</Button>
                </PaginationItem>
                {
                    showStartEllipsis &&
                    <PaginationItem>
                        <PaginationEllipsis />
                    </PaginationItem>
                }
                {
                    currButtons.map((val) => (
                        <PaginationItem key={val + 1}>
                            <Button variant={currentPage === val ? "outline" : "ghost"} onClick={() => handleOnPageClick(val)}>{val}</Button>
                        </PaginationItem>
                    ))
                }
                {
                    showEndEllipsis &&
                    <PaginationItem>
                        <PaginationEllipsis />
                    </PaginationItem>
                }
                <PaginationItem>
                    <Button variant="ghost" onClick={() => handleOnPageClick(currentPage + 1)} disabled={currentPage >= totalPages}>Далее</Button>
                </PaginationItem>
            </PaginationContent>
        </Pagination>
    );
}