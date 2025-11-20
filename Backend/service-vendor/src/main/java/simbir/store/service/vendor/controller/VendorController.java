package simbir.store.service.vendor.controller;


import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.multipart.MultipartFile;
import simbir.store.service.vendor.dto.NoticesDto;
import simbir.store.service.vendor.dto.ProductDto;
import simbir.store.service.vendor.dto.VendorRegistrationDto;
import simbir.store.service.vendor.service.ServiceVendor;

import java.math.BigDecimal;
import java.time.Instant;
import java.util.List;

@RestController
@RequestMapping("/api/vendor")
public class VendorController {

    private final ServiceVendor serviceVendor;

    public VendorController(ServiceVendor serviceVendor) {
        this.serviceVendor = serviceVendor;
    }

    @PostMapping("/register")
    public ResponseEntity<Long> register(@RequestBody VendorRegistrationDto dto) {
        long id = serviceVendor.createVendorRequest(dto);
        return ResponseEntity.ok(id);
    }

    @PostMapping("/{vendorId}/containers/upload")
    public ResponseEntity<Long> uploadContainer(@PathVariable long vendorId,
                                                @RequestPart("file") MultipartFile file) {
        long reqId = serviceVendor.uploadProductContainer(vendorId, file, file.getOriginalFilename());
        return ResponseEntity.ok(reqId);
    }

    @GetMapping("/{vendorId}/products")
    public ResponseEntity<List<ProductDto>> products(@PathVariable long vendorId) {
        return ResponseEntity.ok(serviceVendor.getVendorProducts(vendorId));
    }

    @PutMapping("/{vendorId}/products/{productId}/price")
    public ResponseEntity<ProductDto> changePrice(@PathVariable long vendorId, @PathVariable long productId,
                                                  @RequestParam BigDecimal price) {
        return ResponseEntity.ok(serviceVendor.updateProductPrice(vendorId, productId, price));
    }

    @DeleteMapping("/{vendorId}/products/{productId}")
    public ResponseEntity<Void> deleteProduct(@PathVariable long vendorId, @PathVariable long productId) {
        serviceVendor.deleteProduct(vendorId, productId);
        return ResponseEntity.noContent().build();
    }

    @GetMapping("/{vendorId}/notifications")
    public ResponseEntity<List<NoticesDto>> notifications(@PathVariable long vendorId,
                                                          @RequestParam(required = false) Instant from,
                                                          @RequestParam(required = false) Instant to,
                                                          @RequestParam(defaultValue = "50") int limit) {
        return ResponseEntity.ok(serviceVendor.getNotifications(vendorId, from, to, limit));
    }
}
