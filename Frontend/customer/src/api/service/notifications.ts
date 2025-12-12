import { ErrorDTO, type ErrorDTOType } from "../dtos/base";
import { NotificationsDTO, type NotificationsDTOType } from "../dtos/notifications";

const BASE_URL = `${import.meta.env.VITE_SERVICE_CUSTOMER_URL}/notifications`;
const BASE_PAGE_SIZE = 20;

export async function getNotifications(
    customerId: number,
    page: number,
    from: Date | undefined,
    to: Date | undefined,
) : Promise<NotificationsDTOType | ErrorDTOType> {
    const params = new URLSearchParams(`customerId=${customerId}&page=${page}&size=${BASE_PAGE_SIZE}`);
    if (from !== undefined) {
        params.append("from", from.toISOString());
    }
    if (to !== undefined) {
        params.append("to", to.toISOString());
    }

    const response = await fetch(`${BASE_URL}/${customerId}?${params.toString()}`);
    if (!response.ok) {
        return ErrorDTO.parse({
            message: await response.text(),
        });
    }

    const res = NotificationsDTO.safeParse(await response.json());
    if (!res.success) {
        return ErrorDTO.parse({
            message: res.error.message,
        });
    }
    return res.data;
}