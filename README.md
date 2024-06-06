```
❯ buf --version
1.32.2
```

To reproduce, run `buf generate`.
This will produce:

```
❯ buf generate
protoc: common/v1/value.proto, keyvalue/v1/service.proto
```

**Expected output**:

```
protoc: common/v1/value.proto
protoc: keyvalue/v1/service.proto
```

---

Dropping `include_imports` and adding the import to inputs manually fixes it.

```diff
diff --git a/buf.gen.yaml b/buf.gen.yaml
index 579a6dd..9b4667d 100644
--- a/buf.gen.yaml
+++ b/buf.gen.yaml
@@ -10,10 +10,10 @@ plugins:
   - local: ["go", "run", "./cmd/protoc-gen-custom"]
     out: gen/go
     opt: paths=source_relative
-    include_imports: true
     strategy: directory
 
 inputs:
   - directory: proto
     paths:
       - proto/keyvalue
+      - proto/common
```

```
❯ buf generate
protoc: keyvalue/v1/service.proto
protoc: common/v1/value.proto
```
