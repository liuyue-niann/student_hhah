package com.cxcy.chuangye.domain.dto;


import lombok.Data;

@Data
public class Word {
    private int begin_time;
    private int end_time;
    private String text;
    private String punctuation;
    private boolean fixed;
    
    // getters and setters
}
