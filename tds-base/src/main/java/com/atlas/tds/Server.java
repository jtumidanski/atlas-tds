package com.atlas.tds;

import java.net.URI;

import com.atlas.shared.rest.RestServerFactory;
import com.atlas.shared.rest.UriBuilder;
import com.atlas.tds.constant.RestConstants;

public class Server {
   public static void main(String[] args) {
      URI uri = UriBuilder.host(RestConstants.SERVICE).uri();
      RestServerFactory.create(uri, "com.atlas.tds.rest");
   }
}
