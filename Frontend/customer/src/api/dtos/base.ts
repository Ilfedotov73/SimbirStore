import * as z from "zod"; 

export const ErrorDTO = z.object({
    message: z.string(),
});

export type ErrorDTOType = z.infer<typeof ErrorDTO>;