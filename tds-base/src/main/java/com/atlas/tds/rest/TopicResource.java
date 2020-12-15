package com.atlas.tds.rest;

import com.atlas.tds.rest.processor.TopicProcessor;

import javax.ws.rs.Consumes;
import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.PathParam;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;

@Path("topics")
public class TopicResource {
   @GET
   @Path("")
   @Consumes(MediaType.APPLICATION_JSON)
   @Produces(MediaType.APPLICATION_JSON)
   public Response getTopics() {
      return TopicProcessor.getTopics().build();
   }

   @GET
   @Path("/{id}")
   @Consumes(MediaType.APPLICATION_JSON)
   @Produces(MediaType.APPLICATION_JSON)
   public Response getTopicById(@PathParam("id") String id) {
      return TopicProcessor.getTopic(id).build();
   }
}
