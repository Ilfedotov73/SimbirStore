import { ProductsDTO, ProductDTO, type ProductsDTOType, type ProductDTOType, type ReviewsDTOType, ReviewsDTO, type OfferCreateDTOType } from "../dtos/products";

const BASE_URL = `${import.meta.env.VITE_SERVICE_CUSTOMER_URL}/products`;
const BASE_PAGE_SIZE = 20;

export type ProductSortType = {
    field: string,
    order: string,
};

export async function getProducts(
    page: number, 
    minPrice: number | undefined = undefined, 
    maxPrice: number | undefined = undefined,
    vendorId: number | undefined = undefined,
    query: string | undefined = undefined,
    sort: ProductSortType | undefined = undefined,
) : Promise<ProductsDTOType | Error> {
    const params = new URLSearchParams(`page=${page}&size=${BASE_PAGE_SIZE}`);
    if (minPrice !== undefined) {
        params.append("minPrice", minPrice.toString());
    }
    if (maxPrice !== undefined) {
        params.append("maxPrice", maxPrice.toString());
    }
    if (vendorId !== undefined) {
        params.append("vendorId", vendorId.toString());
    }
    if (query !== undefined) {
        params.append("query", query);
    }
    if (sort !== undefined) {
        params.append("sort", `${sort.field}.${sort.order}`);
    }
    const response = await fetch(`${BASE_URL}?${params.toString()}`);

    if (!response.ok) {
        return new Error(await response.text());
    }

    const res = ProductsDTO.safeParse(await response.json());
    if (!res.success) {
        return new Error(res.error.message);
    }
    return res.data;
}

export async function getProduct(productId: number) : Promise<ProductDTOType | Error> {
    const response = await fetch(`${BASE_URL}/${productId}`);
    if (!response.ok) {
        return new Error(await response.text());
    }

    const res = ProductDTO.safeParse(await response.json());
    if (!res.success) {
        return new Error(res.error.message);
    }
    return res.data;
}

export async function getProductReviews(
    productId: number,
    page: number,
) : Promise<ReviewsDTOType | Error> {
    const response = await fetch(`${BASE_URL}/${productId}/reviews?page=${page}&size=${BASE_PAGE_SIZE}`);
    if (!response.ok) {
        return new Error(await response.text());
    }

    const res = ReviewsDTO.safeParse(await response.json());
    if (!res.success) {
        return new Error(res.error.message);
    }
    return res.data;
}

export async function createOffer(
    productId: number,
    requestBody: OfferCreateDTOType,
) : Promise<null | Error> {
    const response = await fetch(`${BASE_URL}/${productId}/offers`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(requestBody),
    });
    if (!response.ok) {
        return new Error(await response.text());
    }
    return null;
}