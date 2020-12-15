package com.atlas.tds.rest.processor;

import builder.ResultBuilder;
import com.app.rest.util.stream.Collectors;
import com.atlas.tds.ConfigurationRegistry;
import com.atlas.tds.rest.ResultObjectFactory;

public final class TopicProcessor {
   private TopicProcessor() {
   }

   public static ResultBuilder getTopics() {
      return ConfigurationRegistry.getInstance().getConfiguration().topics.stream()
            .map(ResultObjectFactory::create)
            .collect(Collectors.toResultBuilder());
   }

   public static ResultBuilder getTopic(String id) {
      return ConfigurationRegistry.getInstance().getConfiguration().topics.stream()
            .filter(configuration -> configuration.id.equals(id))
            .map(ResultObjectFactory::create)
            .collect(Collectors.toResultBuilder());
   }
}
