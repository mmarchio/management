CREATE OR REPLACE FUNCTION remove_deleted_id_references()
RETURNS TRIGGER AS $$
DECLARE
    deleted_id UUID := OLD.id;
BEGIN
    -- Remove instances of the deleted id from content and manifest columns

    -- Update content column for arrays with suffix _array_model
    UPDATE content
    SET content = jsonb_set(
        content,
        ARRAY[key],
        (content->key - deleted_id::text)::jsonb
    )
    FROM (
        SELECT key
        FROM jsonb_each(content)
        WHERE key LIKE '%_array_model'
    ) AS keys
    WHERE content ?| array_agg(keys.key) -- Check if any of the keys exist in the content
    AND (content->key - deleted_id::text) <> '[]'::jsonb; -- Ensure the array is not empty after removal

    -- Update manifest column for arrays
    UPDATE content
    SET manifest = (manifest - deleted_id::text)::jsonb
    WHERE manifest ? deleted_id::text; -- Check if the id exists in the manifest

    -- Update content column for fields with suffix _model
    UPDATE content
    SET content = jsonb_set(
        content,
        ARRAY[key],
        '""'::jsonb
    )
    FROM (
        SELECT key
        FROM jsonb_each(content)
        WHERE key LIKE '%_model'
          AND content->>key = deleted_id::text
    ) AS keys
    WHERE content ?| array_agg(keys.key); -- Check if any of the keys exist in the content

    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

-- Create a trigger to call this function after delete on content table
CREATE TRIGGER remove_deleted_id_references_trigger
AFTER DELETE ON content
FOR EACH ROW EXECUTE FUNCTION remove_deleted_id_references();