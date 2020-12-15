package com.atlas.tds.rest;

import builder.ResultObjectBuilder;
import com.atlas.tds.configuration.TopicConfiguration;
import com.atlas.tds.rest.attribute.TopicAttributes;
import com.atlas.tds.rest.builder.TopicAttributesBuilder;

public final class ResultObjectFactory {
   private ResultObjectFactory() {
   }

   public static ResultObjectBuilder create(TopicConfiguration configuration) {
      return new ResultObjectBuilder(TopicAttributes.class, configuration.id)
            .setAttribute(new TopicAttributesBuilder().setName(configuration.name));
   }
}
