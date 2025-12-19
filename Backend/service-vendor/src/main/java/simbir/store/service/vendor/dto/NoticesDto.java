package simbir.store.service.vendor.dto;

import lombok.Getter;
import lombok.Setter;

import java.sql.Date;

@Getter
@Setter
public class NoticesDto {
    private int id;

    public int getId() {
        return id;
    }

    public Date getCreateAt() {
        return createAt;
    }

    public String getText() {
        return text;
    }

    public int getEntityId() {
        return entityId;
    }

    private String text;

    public void setCreateAt(Date createAt) {
        this.createAt = createAt;
    }

    public void setEntityId(int entityId) {
        this.entityId = entityId;
    }

    public void setText(String text) {
        this.text = text;
    }

    public void setId(int id) {
        this.id = id;
    }

    private int entityId;
    private Date createAt;
}
