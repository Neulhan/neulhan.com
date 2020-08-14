# Migration `20200815024103-delete-post-field-published`

This migration has been generated by neulhan at 8/15/2020, 2:41:03 AM.
You can check out the [state of the schema](./schema.prisma) after the migration.

## Database Steps

```sql
ALTER TABLE `neulhanDB`.`Post` DROP COLUMN `published`;
```

## Changes

```diff
diff --git schema.prisma schema.prisma
migration 20200815014911-f..20200815024103-delete-post-field-published
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
@@ -25,9 +25,8 @@
 model Post {
     id        String   @default(cuid()) @id
     createdAt DateTime @default(now())
     updatedAt DateTime @updatedAt
-    published Boolean
     title     String
     content   String?
     author   User   @relation(fields: [authorID], references: [id])
```

