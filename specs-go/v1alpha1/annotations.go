package v1alpha1

// The annotations defined in the file can be used for UOR compatibility with
// artifact and image manifests.

const (
	// AnnotationSchema is the reference to the
	// default schema of the collection.
	AnnotationSchema = "uor.schema"
	// AnnotationCollectionLinks references the collections
	// that are linked to a collection node. They will only
	// reference adjacent collection and will not descend
	// into sub-collections.
	AnnotationCollectionLinks = "uor.collections.linked"
	// AnnotationUORAttributes references the collection attributes in a
	// JSON format.
	AnnotationUORAttributes = "uor.attributes"
)
