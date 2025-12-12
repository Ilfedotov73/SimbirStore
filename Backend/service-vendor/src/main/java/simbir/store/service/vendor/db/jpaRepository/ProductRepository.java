package simbir.store.service.vendor.db.jpaRepository;

import org.springframework.data.jpa.repository.JpaRepository;
import simbir.store.service.vendor.db.entity.Product;
import simbir.store.service.vendor.db.entity.User;

import java.util.List;

public interface ProductRepository extends JpaRepository<Product, Long> {

    List<Product> findByVendor(User vendor);

    List<Product> findByNameContainingIgnoreCase(String name);
}
