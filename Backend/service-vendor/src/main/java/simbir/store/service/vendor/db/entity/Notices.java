package simbir.store.service.vendor.db.entity;

import jakarta.persistence.*;
import lombok.Getter;
import lombok.Setter;

import java.util.Date;

@Entity
@Table(name = "notices")
@Getter
@Setter
public class Notices {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Integer id;

    private String text;

    /** FK → User.Id (EntityId) */
    @ManyToOne
    @JoinColumn(name = "entityId")
    private User user;

    private Date createAt;
}
