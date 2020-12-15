package com.atlas.tds;

import java.net.URI;

import com.atlas.shared.rest.RestServerFactory;
import com.atlas.shared.rest.RestService;
import com.atlas.shared.rest.UriBuilder;
import org.glassfish.grizzly.http.server.HttpServer;

import database.PersistenceManager;

public class Server {
   public static void main(String[] args) {
      URI uri = UriBuilder.host(RestService.TOPIC_DISCOVERY).uri();
      RestServerFactory.create(uri, "com.atlas.tds.rest");
   }
}
