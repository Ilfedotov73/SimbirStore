package simbir.store.service.vendor.service;

import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import simbir.store.service.vendor.db.entity.Product;
import simbir.store.service.vendor.db.entity.User;
import simbir.store.service.vendor.db.entity.VendorsProducts;
import simbir.store.service.vendor.db.jpaRepository.ProductRepository;
import simbir.store.service.vendor.db.jpaRepository.UserRepository;
import simbir.store.service.vendor.db.jpaRepository.VendorsProductsRepository;
import simbir.store.service.vendor.dto.ProductContainer;

import java.util.List;

@Service
@RequiredArgsConstructor
public class ProductService {

    private final ProductRepository productRepository;
    private final UserRepository userRepository;
    private final VendorsProductsRepository vendorsProductsRepository;

    /**
     * Создание товара после одобрения очередью.
     * 1) Создаём Product
     * 2) Создаём VendorsProducts (связь)
     */
    public Product createFromXml(ProductContainer container, Long userId) {

        User vendor = userRepository.findById(userId)
                .orElseThrow(() -> new RuntimeException("User not found"));

        if (!"VENDOR".equals(vendor.getRoles())) {
            throw new RuntimeException("User is not a vendor");
        }

        Product p = new Product();
        p.setName(container.getName());
        p.setCharacteristics(container.getCharacteristics());
        p.setPrice(container.getPrice());
        p.setCreateAt(new java.sql.Date(System.currentTimeMillis()));

        Product saved = productRepository.save(p);

        // Создаём связь продавца и товара
        VendorsProducts vp = new VendorsProducts();
        vp.setProduct(saved);
        vp.setVendor(vendor);

        vendorsProductsRepository.save(vp);

        return saved;
    }

    /**
     * Получение всех товаров продавца через таблицу vendors_products.
     */
    public List<Product> getVendorProducts(Long vendorId) {
        return vendorsProductsRepository.findByVendorId(vendorId)
                .stream()
                .map(VendorsProducts::getProduct)
                .toList();
    }

    /**
     * Проверяем принадлежность товара продавцу через связующую таблицу.
     */
    private VendorsProducts getLinkOrThrow(Long productId, Long vendorId) {
        return vendorsProductsRepository.findByVendorIdAndProductId(vendorId, productId)
                .orElseThrow(() ->
                        new RuntimeException("Product does not belong to this vendor"));
    }

    /**
     * Редактирование товара.
     */
    public Product updateProduct(Long productId, Long vendorId, Float price, String characteristics) {

        // Проверка принадлежности
        getLinkOrThrow(productId, vendorId);

        Product p = productRepository.findById(Long.valueOf(productId))
                .orElseThrow();

        if (price != null) p.setPrice(price);
        if (characteristics != null) p.setCharacteristics(characteristics);

        return productRepository.save(p);
    }

    /**
     * Удаление товара продавцом.
     */
    public void deleteProduct(Long productId, Long vendorId) {

        VendorsProducts link = getLinkOrThrow(productId, vendorId);

        // Сначала удалить связь
        vendorsProductsRepository.delete(link);

        // Потом удалить сам продукт
        productRepository.deleteById(productId);
    }
}
