# object

Config contains the configuration for the openlane server


**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**refreshInterval**|`integer`|RefreshInterval determines how often to reload the config<br/>||
|[**server**](#server)|`object`|Server settings for the echo server<br/>|yes|
|[**db**](#db)|`object`||yes|
|[**providers**](#providers)|`object`|||
|[**turso**](#turso)|`object`||yes|
|[**redis**](#redis)|`object`|||
|[**tracer**](#tracer)|`object`|||
|[**sessions**](#sessions)|`object`|||
|[**ratelimit**](#ratelimit)|`object`|||

**Additional Properties:** not allowed  
<a name="server"></a>
## server: object

Server settings for the echo server


**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**debug**|`boolean`|Debug enables debug mode for the server<br/>|no|
|**dev**|`boolean`|Dev enables echo's dev mode options<br/>|no|
|**listen**|`string`|Listen sets the listen address to serve the echo server on<br/>|yes|
|**shutdownGracePeriod**|`integer`|ShutdownGracePeriod sets the grace period for in flight requests before shutting down<br/>|no|
|**readTimeout**|`integer`|ReadTimeout sets the maximum duration for reading the entire request including the body<br/>|no|
|**writeTimeout**|`integer`|WriteTimeout sets the maximum duration before timing out writes of the response<br/>|no|
|**idleTimeout**|`integer`|IdleTimeout sets the maximum amount of time to wait for the next request when keep-alives are enabled<br/>|no|
|**readHeaderTimeout**|`integer`|ReadHeaderTimeout sets the amount of time allowed to read request headers<br/>|no|
|[**tls**](#servertls)|`object`|TLS settings for the server for secure connections<br/>|no|
|[**cors**](#servercors)|`object`||no|
|[**secure**](#serversecure)|`object`||no|
|[**redirects**](#serverredirects)|`object`||no|
|[**cacheControl**](#servercachecontrol)|`object`||no|
|[**mime**](#servermime)|`object`||no|

**Additional Properties:** not allowed  
<a name="servertls"></a>
### server\.tls: object

TLS settings for the server for secure connections


**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**enabled**|`boolean`|Enabled turns on TLS settings for the server<br/>||
|**certFile**|`string`|CertFile location for the TLS server<br/>||
|**certKey**|`string`|CertKey file location for the TLS server<br/>||
|**autoCert**|`boolean`|AutoCert generates the cert with letsencrypt, this does not work on localhost<br/>||

**Additional Properties:** not allowed  
<a name="servercors"></a>
### server\.cors: object

**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**enabled**|`boolean`|||
|[**prefixes**](#servercorsprefixes)|`object`|||
|[**allowOrigins**](#servercorsalloworigins)|`string[]`|||
|**cookieInsecure**|`boolean`|||

**Additional Properties:** not allowed  
<a name="servercorsprefixes"></a>
#### server\.cors\.prefixes: object

**Additional Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|[**Additional Properties**](#servercorsprefixesadditionalproperties)|`string[]`|||

<a name="servercorsprefixesadditionalproperties"></a>
##### server\.cors\.prefixes\.additionalProperties: array

**Items**

**Item Type:** `string`  
<a name="servercorsalloworigins"></a>
#### server\.cors\.allowOrigins: array

**Items**

**Item Type:** `string`  
<a name="serversecure"></a>
### server\.secure: object

**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**enabled**|`boolean`|||
|**xssprotection**|`string`|||
|**contenttypenosniff**|`string`|||
|**xframeoptions**|`string`|||
|**hstspreloadenabled**|`boolean`|||
|**hstsmaxage**|`integer`|||
|**contentsecuritypolicy**|`string`|||
|**referrerpolicy**|`string`|||
|**cspreportonly**|`boolean`|||

**Additional Properties:** not allowed  
<a name="serverredirects"></a>
### server\.redirects: object

**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**enabled**|`boolean`|||
|[**redirects**](#serverredirectsredirects)|`object`|||
|**code**|`integer`|||

**Additional Properties:** not allowed  
<a name="serverredirectsredirects"></a>
#### server\.redirects\.redirects: object

**Additional Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**Additional Properties**|`string`|||

<a name="servercachecontrol"></a>
### server\.cacheControl: object

**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**enabled**|`boolean`|||
|[**noCacheHeaders**](#servercachecontrolnocacheheaders)|`object`|||
|[**etagHeaders**](#servercachecontroletagheaders)|`string[]`|||

**Additional Properties:** not allowed  
<a name="servercachecontrolnocacheheaders"></a>
#### server\.cacheControl\.noCacheHeaders: object

**Additional Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**Additional Properties**|`string`|||

<a name="servercachecontroletagheaders"></a>
#### server\.cacheControl\.etagHeaders: array

**Items**

**Item Type:** `string`  
<a name="servermime"></a>
### server\.mime: object

**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**enabled**|`boolean`|||
|**mimeTypesFile**|`string`|||
|**defaultContentType**|`string`|||

**Additional Properties:** not allowed  
<a name="db"></a>
## db: object

**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**debug**|`boolean`|debug enables printing the debug database logs<br/>|no|
|**databaseName**|`string`|the name of the database to use with otel tracing<br/>|no|
|**driverName**|`string`|sql driver name<br/>|no|
|**multiWrite**|`boolean`|enables writing to two databases simultaneously<br/>|no|
|**primaryDbSource**|`string`|dsn of the primary database<br/>|yes|
|**secondaryDbSource**|`string`|dsn of the secondary database if multi-write is enabled<br/>|no|
|**cacheTTL**|`integer`|cache results for subsequent requests<br/>|no|
|**runMigrations**|`boolean`|run migrations on startup<br/>|no|
|**migrationProvider**|`string`|migration provider to use for running migrations<br/>|no|
|**enableHistory**|`boolean`|enable history data to be logged to the database<br/>|no|

**Additional Properties:** not allowed  
<a name="providers"></a>
## providers: object

**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**tursoEnabled**|`boolean`|TursoEnabled enables the turso provider<br/>||
|**localEnabled**|`boolean`|LocalEnabled enables the local provider<br/>||

**Additional Properties:** not allowed  
<a name="turso"></a>
## turso: object

**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**token**|`string`||yes|
|**baseUrl**|`string`||yes|
|**orgName**|`string`||yes|

**Additional Properties:** not allowed  
<a name="redis"></a>
## redis: object

**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**enabled**|`boolean`|||
|**address**|`string`|||
|**name**|`string`|||
|**username**|`string`|||
|**password**|`string`|||
|**db**|`integer`|||
|**dialTimeout**|`integer`|||
|**readTimeout**|`integer`|||
|**writeTimeout**|`integer`|||
|**maxRetries**|`integer`|||
|**minIdleConns**|`integer`|||
|**maxIdleConns**|`integer`|||
|**maxActiveConns**|`integer`|||

**Additional Properties:** not allowed  
<a name="tracer"></a>
## tracer: object

**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**enabled**|`boolean`|||
|**provider**|`string`|||
|**environment**|`string`|||
|[**stdout**](#tracerstdout)|`object`|||
|[**otlp**](#tracerotlp)|`object`|||

**Additional Properties:** not allowed  
<a name="tracerstdout"></a>
### tracer\.stdout: object

**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**pretty**|`boolean`|||
|**disableTimestamp**|`boolean`|||

**Additional Properties:** not allowed  
<a name="tracerotlp"></a>
### tracer\.otlp: object

**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**endpoint**|`string`|||
|**insecure**|`boolean`|||
|**certificate**|`string`|||
|[**headers**](#tracerotlpheaders)|`string[]`|||
|**compression**|`string`|||
|**timeout**|`integer`|||

**Additional Properties:** not allowed  
<a name="tracerotlpheaders"></a>
#### tracer\.otlp\.headers: array

**Items**

**Item Type:** `string`  
<a name="sessions"></a>
## sessions: object

**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**signingKey**|`string`|||
|**encryptionKey**|`string`|||
|**domain**|`string`|||
|**maxAge**|`integer`|||

**Additional Properties:** not allowed  
<a name="ratelimit"></a>
## ratelimit: object

**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**enabled**|`boolean`|||
|**limit**|`number`|||
|**burst**|`integer`|||
|**expires**|`integer`|||

**Additional Properties:** not allowed  

