package com.atlas.tds.rest.builder;

import builder.AttributeResultBuilder;
import builder.RecordBuilder;
import com.atlas.tds.rest.attribute.TopicAttributes;

public class TopicAttributesBuilder extends RecordBuilder<TopicAttributes, TopicAttributesBuilder> implements AttributeResultBuilder {
   private static final String NAME = "NAME";

   @Override
   public TopicAttributes construct() {
      return new TopicAttributes(get(NAME));
   }

   @Override
   public TopicAttributesBuilder getThis() {
      return this;
   }

   public TopicAttributesBuilder setName(String name) {
      return set(NAME, name);
   }

}
