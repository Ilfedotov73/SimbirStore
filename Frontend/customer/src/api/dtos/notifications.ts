import * as z from "zod"; 
import { BasePagingDTO } from "./paging";

export const NotificationDTO = z.object({
    id: z.int(),
    text: z.string(),
    entityId: z.int(),
    createAt: z.iso.datetime(),
});

export const NotificationsDTO = BasePagingDTO.extend({
    items: z.array(NotificationDTO)
});

export type NotificationDTOType = z.infer<typeof NotificationDTO>;
export type NotificationsDTOType = z.infer<typeof NotificationsDTO>;