package com.cxcy.chuangye.domain.dto;

import lombok.Data;

import java.util.List;

@Data
public class Stash {
    private int sentence_id;
    private int begin_time;
    private String text;
    private List<?> words;
    
    // getters and setters
}
