import * as z from "zod"; 

export const BasePagingDTO = z.object({
    paging: z.object({
        page: z.int(),
        size: z.int(),
        total: z.int()
    })
});

export type BasePagingType = z.infer<typeof BasePagingDTO>;