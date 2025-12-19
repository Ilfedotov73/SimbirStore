package simbir.store.service.vendor.db.entity;

import jakarta.persistence.*;
import lombok.Getter;
import lombok.Setter;

@Entity
@Table(name = "vendors_products")
@Getter
@Setter
public class VendorsProducts {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Integer id;

    public User getVendor() {
        return vendor;
    }

    public void setId(Integer id) {
        this.id = id;
    }

    public void setProduct(Product product) {
        this.product = product;
    }

    public void setVendor(User vendor) {
        this.vendor = vendor;
    }

    public Integer getId() {
        return id;
    }

    public Product getProduct() {
        return product;
    }

    /** FK → User.Id */
    @ManyToOne
    @JoinColumn(name = "user_id")
    private User vendor;

    /** FK → Products.Id */
    @ManyToOne
    @JoinColumn(name = "product_id")
    private Product product;
}
