import { ErrorDTO, type ErrorDTOType } from "../dtos/base";
import { CustomerDTO, type CustomerDTOType } from "../dtos/customers";

const BASE_URL = `${import.meta.env.VITE_SERVICE_CUSTOMER_URL}/customers`;

export async function getCustomer(customerId: number) : Promise<CustomerDTOType | ErrorDTOType> {
    const response = await fetch(`${BASE_URL}/${customerId}`);
    if (!response.ok) {
        return ErrorDTO.parse({
            message: await response.text(),
        });
    }

    const res = CustomerDTO.safeParse(await response.json());
    if (!res.success) {
        return ErrorDTO.parse({
            message: res.error.message,
        });
    }
    return res.data;
}