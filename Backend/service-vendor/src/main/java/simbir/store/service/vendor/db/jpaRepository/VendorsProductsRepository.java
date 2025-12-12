package simbir.store.service.vendor.db.jpaRepository;

import org.springframework.data.jpa.repository.JpaRepository;
import simbir.store.service.vendor.db.entity.VendorsProducts;

import java.util.List;
import java.util.Optional;

public interface VendorsProductsRepository extends JpaRepository<VendorsProducts, Long> {

    List<VendorsProducts> findByVendorId(Long vendorId);

    Optional<VendorsProducts> findByVendorIdAndProductId(Long vendorId, Long productId);
}
