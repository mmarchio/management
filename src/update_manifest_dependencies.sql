CREATE OR REPLACE FUNCTION update_manifest()
RETURNS TRIGGER AS $$
DECLARE
    target_id UUID := NEW.id;
BEGIN
    -- CTE to gather all dependent ids using recursion
    WITH RECURSIVE dependent_ids AS (
        -- Initial selection: start with the given id
        SELECT id, content
        FROM content
        WHERE id = target_id

        UNION ALL

        -- Recursive selection: find records whose id is in the content of the previous records
        SELECT c.id, c.content
        FROM content c
        JOIN dependent_ids d ON 
            (
                c.id::text = ANY(
                    jsonb_array_elements_text(
                        jsonb_extract_path(
                            d.content, key
                        ) FILTER (
                            WHERE key LIKE '%_model'
                        )
                    ) OR c.id::text = ANY(
                    ARRAY_AGG(key) FILTER (WHERE key LIKE '%_array_model')
                    )
                )
            )
        WHERE c.id <> d.id
    ),
    -- Collect all distinct ids from the recursive CTE
    collected_ids AS (
        SELECT DISTINCT id AS dependent_id
        FROM dependent_ids
    )
    -- Update the manifest column of the queried record with the collected dependent ids
    UPDATE content
    SET manifest = (SELECT jsonb_agg(dependent_id::text) FROM collected_ids)
    WHERE id = target_id;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create a trigger to call this function on insert or update
CREATE TRIGGER update_manifest_trigger
AFTER INSERT OR UPDATE ON content
FOR EACH ROW EXECUTE FUNCTION update_manifest();

CREATE OR REPLACE FUNCTION update_manifest_with_dependencies(target_id UUID)
RETURNS VOID AS
$$
WITH RECURSIVE dependency_cte AS (
    -- Initial selection: start with the given id
    SELECT id, content
    FROM content
    WHERE id = target_id

    UNION ALL

    -- Recursive selection: find records whose id is in the content of the previous records
    SELECT c.id, c.content
    FROM content c
    JOIN dependent_ids d ON 
        c.id::text = ANY(
            jsonb_array_elements_text(jsonb_extract_path(d.content, 'nodes')) || 
            array_agg(key)
        )
    WHERE c.id <> d.id
),
-- Collect all distinct ids from the recursive CTE
collected_ids AS (
    SELECT DISTINCT id AS dependent_id
    FROM dependent_ids
),
-- Extract keys matching the pattern %_model and their values
matching_keys AS (
    SELECT key, value::text AS uuid_value
    FROM content, jsonb_each(content)
    WHERE id = target_id
      AND key LIKE '%_model'
),
-- Combine initial ids with matching keys
combined_ids AS (
    SELECT DISTINCT dependent_id
    FROM collected_ids

    UNION ALL

    SELECT uuid_value::uuid
    FROM matching_keys
)
-- Update the manifest column of the queried record with the combined dependent ids
UPDATE content
SET manifest = (
    SELECT jsonb_agg(dependent_id::text) 
    FROM combined_ids
)
WHERE id = target_id;
$$ LANGUAGE SQL;