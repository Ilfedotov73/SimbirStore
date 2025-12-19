package simbir.store.service.vendor.db.entity;

import jakarta.persistence.*;
import lombok.Getter;
import lombok.Setter;

import java.util.Date;
import java.util.List;

@Entity
@Table(name = "products")
@Getter
@Setter
public class Product {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Integer id;

    private String name;
    private Float price;

    private String photoUrl;

    public void setId(Integer id) {
        this.id = id;
    }

    public void setName(String name) {
        this.name = name;
    }

    public void setPrice(Float price) {
        this.price = price;
    }

    public void setPhotoUrl(String photoUrl) {
        this.photoUrl = photoUrl;
    }

    public void setCreateAt(Date createAt) {
        this.createAt = createAt;
    }

    public void setCharacteristics(String characteristics) {
        this.characteristics = characteristics;
    }

    public void setProductRating(Float productRating) {
        this.productRating = productRating;
    }

    public void setVendorLinks(List<VendorsProducts> vendorLinks) {
        this.vendorLinks = vendorLinks;
    }

    public void setReviews(List<ProductsReviews> reviews) {
        this.reviews = reviews;
    }

    public Integer getId() {
        return id;
    }

    public String getName() {
        return name;
    }

    public Float getPrice() {
        return price;
    }

    public String getPhotoUrl() {
        return photoUrl;
    }

    public Date getCreateAt() {
        return createAt;
    }

    public Float getProductRating() {
        return productRating;
    }

    public String getCharacteristics() {
        return characteristics;
    }

    public List<VendorsProducts> getVendorLinks() {
        return vendorLinks;
    }

    public List<ProductsReviews> getReviews() {
        return reviews;
    }

    private Date createAt;

    private String characteristics;

    private Float productRating;

    /** связь с VendorsProducts (m:m через entity) */
    @OneToMany(mappedBy = "product")
    private List<VendorsProducts> vendorLinks;

    /** список отзывов (m:m через ProductsReviews) */
    @OneToMany(mappedBy = "product")
    private List<ProductsReviews> reviews;
}
