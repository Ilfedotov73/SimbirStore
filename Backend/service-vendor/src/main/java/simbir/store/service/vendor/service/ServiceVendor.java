package simbir.store.service.vendor.service;

import org.springframework.web.multipart.MultipartFile;
import simbir.store.service.vendor.dto.NoticesDto;
import simbir.store.service.vendor.dto.ProductDto;
import simbir.store.service.vendor.dto.VendorDto;
import simbir.store.service.vendor.dto.VendorRegistrationDto;

import java.math.BigDecimal;
import java.time.Instant;
import java.util.List;

public interface ServiceVendor {

    // Регистрация продавца (создание заявки на рассмотрение)
    long createVendorRequest(VendorRegistrationDto dto);

    VendorDto getVendorRequest(long requestId);

    VendorDto getVendorById(long vendorId);

    // Обработка заявки (вызывается TaskManager/админ)
    VendorDto processVendorRequest(long requestId, boolean approve, String adminComment);

    // Контейнер/заявка на загрузку товаров (upload xml)
    long uploadProductContainer(long vendorId, MultipartFile xmlFile, String fileName);

    List<ProductDto> getVendorProducts(long vendorId);
    ProductDto updateProductPrice(long vendorId, long productId, BigDecimal newPrice);

    void deleteProduct(long vendorId, long productId);

    List<NoticesDto> getNotifications(long vendorId, Instant from, Instant to, int limit);
}
