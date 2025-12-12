package simbir.store.service.vendor.dto;

import jakarta.xml.bind.annotation.XmlRootElement;
import lombok.Getter;
import lombok.Setter;

@XmlRootElement(name = "ProductContainer")
@Getter
@Setter
public class ProductContainer {
    private String name;
    private Float price;
    private String characteristics;
}
