package simbir.store.service.vendor.db.jpaRepository;

import org.springframework.data.jpa.repository.JpaRepository;
import simbir.store.service.vendor.db.entity.ProductsReviews;

import java.util.List;

public interface ProductsReviewsRepository extends JpaRepository<ProductsReviews, Integer> {

    List<ProductsReviews> findByProductId(Long productId);

    List<ProductsReviews> findByCustomerId(Long userId);
}
