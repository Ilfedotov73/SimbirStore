import * as z from "zod"; 

export const VendorDTO = z.object({
    id: z.int(),
    firstName: z.string(),
    lastName: z.string(),
    phoneNumber: z.string(),
    photoUrl: z.url(),
    vendorTelegramId: z.string(),
    createAt: z.iso.datetime()
});

export type VendorDTOType = z.infer<typeof VendorDTO>;