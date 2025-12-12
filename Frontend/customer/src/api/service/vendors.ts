import { ErrorDTO, type ErrorDTOType } from "../dtos/base";
import { ProductsDTO, type ProductsDTOType } from "../dtos/products";
import { VendorDTO, type VendorDTOType } from "../dtos/vendors";

const BASE_URL = `${import.meta.env.SERVICE_CUSTOMER_URL}/vendors`;
const BASE_PAGE_SIZE = 20;

export async function getVendor(vendorId: number) : Promise<VendorDTOType | ErrorDTOType> {
    const response = await fetch(`${BASE_URL}/${vendorId}`);
    if (!response.ok) {
        return ErrorDTO.parse({
            message: await response.text(),
        });
    }

    const res = VendorDTO.safeParse(await response.json());
    if (!res.success) {
        return ErrorDTO.parse({
            message: res.error.message,
        });
    }
    return res.data;
}

export async function getVendorProducts(vendorId: number, page: number) : Promise<ProductsDTOType | ErrorDTOType> {
    const response = await fetch(`${BASE_URL}/${vendorId}?page=${page}&size=${BASE_PAGE_SIZE}`);
    if (!response.ok) {
        return ErrorDTO.parse({
            message: await response.text(),
        });
    }

    const res = ProductsDTO.safeParse(await response.json());
    if (!res.success) {
        return ErrorDTO.parse({
            message: res.error.message,
        });
    }
    return res.data;
}