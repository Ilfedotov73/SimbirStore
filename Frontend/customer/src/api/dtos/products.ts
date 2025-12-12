import * as z from "zod"; 
import { BasePagingDTO } from "./paging";
import { VendorDTO } from "./vendors";

export const ProductShortDTO = z.object({
    id: z.int(),
    name: z.string(),
    price: z.number(),
    photoUrl: z.url(),
    characteristics: z.string(),
    productRating: z.number(),
    vendorId: z.int()
});

export const ProductDTO = z.object({
    product: ProductShortDTO.extend({
        createAt: z.iso.datetime(),
    }),
    vendor: VendorDTO,
    reviewsCount: z.int(),
    averageRating: z.number(),
});

export const ProductsDTO = BasePagingDTO.extend({
    items: z.array(ProductShortDTO)
});

export const ReviewDTO = z.object({
    id: z.int(),
    customerId: z.int(),
    productId: z.int(),
    review: z.string(),
    rating: z.int(),
});

export const ReviewsDTO = BasePagingDTO.extend({
    items: z.array(ReviewDTO)
});

export const OfferCreateDTO = z.object({
    buyerId: z.int(),
    vendorId: z.int(),
    offerPrice: z.number(),
    message: z.string(),
});

export type ProductShortDTOType = z.infer<typeof ProductShortDTO>;
export type ProductDTOType = z.infer<typeof ProductDTO>;
export type ProductsDTOType = z.infer<typeof ProductsDTO>;

export type ReviewDTOType = z.infer<typeof ReviewDTO>;
export type ReviewsDTOType = z.infer<typeof ReviewsDTO>;

export type OfferCreateDTOType = z.infer<typeof OfferCreateDTO>;