import * as z from "zod"; 

export const CustomerDTO = z.object({
    id: z.int(),
    firstName: z.string(),
    lastName: z.string(),
    phoneNumber: z.string(),
    photoUrl: z.url(),
    customerTelegramId: z.string(),
    createAt: z.iso.datetime(),
    login: z.string(),
    email: z.email(),
});

export type CustomerDTOType = z.infer<typeof CustomerDTO>;