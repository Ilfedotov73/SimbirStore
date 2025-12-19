package simbir.store.service.vendor.dto;

import jakarta.xml.bind.annotation.XmlRootElement;
import lombok.Getter;
import lombok.Setter;

@XmlRootElement(name = "ProductContainer")
@Getter
@Setter
public class ProductContainer {
    private String name;

    public Float getPrice() {
        return price;
    }

    public void setName(String name) {
        this.name = name;
    }

    public void setPrice(Float price) {
        this.price = price;
    }

    public void setCharacteristics(String characteristics) {
        this.characteristics = characteristics;
    }

    public String getName() {
        return name;
    }

    public String getCharacteristics() {
        return characteristics;
    }

    private Float price;
    private String characteristics;
}
