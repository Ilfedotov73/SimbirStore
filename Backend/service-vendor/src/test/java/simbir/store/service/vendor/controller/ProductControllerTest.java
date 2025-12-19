package simbir.store.service.vendor.controller;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.MockitoAnnotations;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.setup.MockMvcBuilders;
import simbir.store.service.vendor.db.entity.Product;
import simbir.store.service.vendor.dto.ProductContainer;
import simbir.store.service.vendor.service.ProductService;

import java.util.List;

import static org.mockito.ArgumentMatchers.*;
import static org.mockito.Mockito.doNothing;
import static org.mockito.Mockito.when;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.*;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.*;

class ProductControllerTest {

    private MockMvc mockMvc;
    private ObjectMapper objectMapper = new ObjectMapper();

    @Mock
    private ProductService productService;

    @InjectMocks
    private ProductController productController;

    @BeforeEach
    void setup() {
        MockitoAnnotations.openMocks(this);
        mockMvc = MockMvcBuilders.standaloneSetup(productController).build();
    }

    @Test
    void testCreateProduct() throws Exception {
        Product product = new Product();
        product.setId(1);
        product.setName("Product1");

        when(productService.createFromXml(any(ProductContainer.class), anyLong()))
                .thenReturn(product);

        ProductContainer container = new ProductContainer();
        container.setName("Product1");

        mockMvc.perform(post("/api/products/vendor/1")
                        .contentType(MediaType.APPLICATION_JSON)
                        .content(objectMapper.writeValueAsString(container)))
                .andExpect(status().isCreated())
                .andExpect(jsonPath("$.id").value(1))
                .andExpect(jsonPath("$.name").value("Product1"));
    }

    @Test
    void testGetVendorProducts() throws Exception {
        Product product = new Product();
        product.setId(1);
        product.setName("Product1");

        when(productService.getVendorProducts(anyLong())).thenReturn(List.of(product));

        mockMvc.perform(get("/api/products/vendor/1"))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$[0].name").value("Product1"));
    }

    @Test
    void testUpdateProduct() throws Exception {
        Product product = new Product();
        product.setId(1);
        product.setPrice(10f);
        product.setCharacteristics("Updated");

        when(productService.updateProduct(anyLong(), anyLong(), any(), any()))
                .thenReturn(product);

        mockMvc.perform(put("/api/products/vendor/1/1")
                        .param("price", "10")
                        .param("characteristics", "Updated"))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.price").value(10))
                .andExpect(jsonPath("$.characteristics").value("Updated"));
    }

    @Test
    void testDeleteProduct() throws Exception {
        doNothing().when(productService).deleteProduct(anyLong(), anyLong());

        mockMvc.perform(delete("/api/products/vendor/1/1"))
                .andExpect(status().isNoContent());
    }
}
