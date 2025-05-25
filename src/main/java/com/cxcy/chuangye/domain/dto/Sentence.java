package com.cxcy.chuangye.domain.dto;

import lombok.Data;

import java.util.List;

@Data
public class Sentence {
    private int begin_time;
    private int end_time;
    private String text;
    private List<Word> words;
    private Stash stash;
    private boolean heartbeat;
    private int sentence_id;
    
    // getters and setters
}

