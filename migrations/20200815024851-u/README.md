# Migration `20200815024851-u`

This migration has been generated by neulhan at 8/15/2020, 2:48:51 AM.
You can check out the [state of the schema](./schema.prisma) after the migration.

## Database Steps

```sql
ALTER TABLE `neulhanDB`.`Post` DROP COLUMN `content`,
ADD COLUMN `content` varchar(191) NOT NULL  ;
```

## Changes

```diff
diff --git schema.prisma schema.prisma
migration 20200815024103-delete-post-field-published..20200815024851-u
--- datamodel.dml
+++ datamodel.dml
@@ -1,8 +1,8 @@
 datasource db {
     provider = "mysql"
-    // url = "***"
-    url = "***"
+    // url      = env("DATABASE_URL")
+    url      = "mysql://neulhan:gksruf@localhost:3306/neulhanDB"
 }
@@ -26,9 +26,9 @@
     id        String   @default(cuid()) @id
     createdAt DateTime @default(now())
     updatedAt DateTime @updatedAt
     title     String
-    content   String?
+    content   String
     author   User   @relation(fields: [authorID], references: [id])
     authorID String
 }
```


