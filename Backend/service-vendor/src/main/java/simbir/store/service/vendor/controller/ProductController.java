package simbir.store.service.vendor.controller;

import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;
import simbir.store.service.vendor.db.entity.Product;
import simbir.store.service.vendor.dto.ProductContainer;
import simbir.store.service.vendor.service.ProductService;

import java.util.List;

@RestController
@RequestMapping("/api/products")
@RequiredArgsConstructor
public class ProductController {

    private final ProductService productService;

    /**
     * Создание товара для продавца
     */
    @PostMapping("/vendor/{vendorId}")
    @ResponseStatus(HttpStatus.CREATED)
    public Product createProduct(@PathVariable Long vendorId, @RequestBody ProductContainer container) {
        return productService.createFromXml(container, vendorId);
    }

    /**
     * Получение всех товаров конкретного продавца
     */
    @GetMapping("/vendor/{vendorId}")
    public List<Product> getVendorProducts(@PathVariable Long vendorId) {
        return productService.getVendorProducts(vendorId);
    }

    /**
     * Редактирование товара продавцом
     */
    @PutMapping("/vendor/{vendorId}/{productId}")
    public Product updateProduct(
            @PathVariable Long vendorId,
            @PathVariable Long productId,
            @RequestParam(required = false) Float price,
            @RequestParam(required = false) String characteristics
    ) {
        return productService.updateProduct(productId, vendorId, price, characteristics);
    }

    /**
     * Удаление товара продавцом
     */
    @DeleteMapping("/vendor/{vendorId}/{productId}")
    @ResponseStatus(HttpStatus.NO_CONTENT)
    public void deleteProduct(@PathVariable Long vendorId, @PathVariable Long productId) {
        productService.deleteProduct(productId, vendorId);
    }
}
