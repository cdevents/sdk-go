/*
Copyright 2024 The CDEvents Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

SPDX-License-Identifier: Apache-2.0
*/

package api

import "fmt"

type SchemaDB map[string]string

var SchemasById = map[string]string{
	"https://cdevents.dev/0.3.0/schema/artifact-packaged-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/artifact-packaged-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.artifact.packaged.0.1.1"
          ],
          "default": "dev.cdevents.artifact.packaged.0.1.1"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "artifact"
          ],
          "default": "artifact"
        },
        "content": {
          "properties": {
            "change": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "change"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}`,
	"https://cdevents.dev/0.3.0/schema/artifact-published-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/artifact-published-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.artifact.published.0.1.1"
          ],
          "default": "dev.cdevents.artifact.published.0.1.1"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "artifact"
          ],
          "default": "artifact"
        },
        "content": {
          "properties": {},
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}`,
	"https://cdevents.dev/0.3.0/schema/artifact-signed-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/artifact-signed-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.artifact.signed.0.1.0"
          ],
          "default": "dev.cdevents.artifact.signed.0.1.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "artifact"
          ],
          "default": "artifact"
        },
        "content": {
          "properties": {
            "signature": {
              "type": "string",
              "minLength": 1
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "signature"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}`,
	"https://cdevents.dev/0.3.0/schema/branch-created-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/branch-created-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.branch.created.0.1.2"
          ],
          "default": "dev.cdevents.branch.created.0.1.2"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "branch"
          ],
          "default": "branch"
        },
        "content": {
          "properties": {
            "repository": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}`,
	"https://cdevents.dev/0.3.0/schema/branch-deleted-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/branch-deleted-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.branch.deleted.0.1.2"
          ],
          "default": "dev.cdevents.branch.deleted.0.1.2"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "branch"
          ],
          "default": "branch"
        },
        "content": {
          "properties": {
            "repository": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}`,
	"https://cdevents.dev/0.3.0/schema/build-finished-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/build-finished-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.build.finished.0.1.1"
          ],
          "default": "dev.cdevents.build.finished.0.1.1"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "build"
          ],
          "default": "build"
        },
        "content": {
          "properties": {
            "artifactId": {
              "type": "string"
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}`,
	"https://cdevents.dev/0.3.0/schema/build-queued-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/build-queued-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.build.queued.0.1.1"
          ],
          "default": "dev.cdevents.build.queued.0.1.1"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "build"
          ],
          "default": "build"
        },
        "content": {
          "properties": {},
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}`,
	"https://cdevents.dev/0.3.0/schema/build-started-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/build-started-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.build.started.0.1.1"
          ],
          "default": "dev.cdevents.build.started.0.1.1"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "build"
          ],
          "default": "build"
        },
        "content": {
          "properties": {},
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}`,
	"https://cdevents.dev/0.3.0/schema/change-abandoned-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/change-abandoned-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.change.abandoned.0.1.2"
          ],
          "default": "dev.cdevents.change.abandoned.0.1.2"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "change"
          ],
          "default": "change"
        },
        "content": {
          "properties": {
            "repository": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}`,
	"https://cdevents.dev/0.3.0/schema/change-created-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/change-created-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.change.created.0.1.2"
          ],
          "default": "dev.cdevents.change.created.0.1.2"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "change"
          ],
          "default": "change"
        },
        "content": {
          "properties": {
            "repository": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}`,
	"https://cdevents.dev/0.3.0/schema/change-merged-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/change-merged-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.change.merged.0.1.2"
          ],
          "default": "dev.cdevents.change.merged.0.1.2"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "change"
          ],
          "default": "change"
        },
        "content": {
          "properties": {
            "repository": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}`,
	"https://cdevents.dev/0.3.0/schema/change-reviewed-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/change-reviewed-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.change.reviewed.0.1.2"
          ],
          "default": "dev.cdevents.change.reviewed.0.1.2"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "change"
          ],
          "default": "change"
        },
        "content": {
          "properties": {
            "repository": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}`,
	"https://cdevents.dev/0.3.0/schema/change-updated-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/change-updated-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.change.updated.0.1.2"
          ],
          "default": "dev.cdevents.change.updated.0.1.2"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "change"
          ],
          "default": "change"
        },
        "content": {
          "properties": {
            "repository": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}`,
	"https://cdevents.dev/0.3.0/schema/environment-created-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/environment-created-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.environment.created.0.1.1"
          ],
          "default": "dev.cdevents.environment.created.0.1.1"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "environment"
          ],
          "default": "environment"
        },
        "content": {
          "properties": {
            "name": {
              "type": "string"
            },
            "url": {
              "type": "string"
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}`,
	"https://cdevents.dev/0.3.0/schema/environment-deleted-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/environment-deleted-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.environment.deleted.0.1.1"
          ],
          "default": "dev.cdevents.environment.deleted.0.1.1"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "environment"
          ],
          "default": "environment"
        },
        "content": {
          "properties": {
            "name": {
              "type": "string"
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}`,
	"https://cdevents.dev/0.3.0/schema/environment-modified-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/environment-modified-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.environment.modified.0.1.1"
          ],
          "default": "dev.cdevents.environment.modified.0.1.1"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "environment"
          ],
          "default": "environment"
        },
        "content": {
          "properties": {
            "name": {
              "type": "string"
            },
            "url": {
              "type": "string"
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}`,
	"https://cdevents.dev/0.3.0/schema/incident-detected-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/incident-detected-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.incident.detected.0.1.0"
          ],
          "default": "dev.cdevents.incident.detected.0.1.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "incident"
          ],
          "default": "incident"
        },
        "content": {
          "properties": {
            "description": {
              "type": "string"
            },
            "environment": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "service": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "artifactId": {
              "type": "string",
              "minLength": 1
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "environment"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}`,
	"https://cdevents.dev/0.3.0/schema/incident-reported-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/incident-reported-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.incident.reported.0.1.0"
          ],
          "default": "dev.cdevents.incident.reported.0.1.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "incident"
          ],
          "default": "incident"
        },
        "content": {
          "properties": {
            "description": {
              "type": "string"
            },
            "environment": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "ticketURI": {
              "type": "string",
              "format": "uri",
              "minLength": 1
            },
            "service": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "artifactId": {
              "type": "string",
              "minLength": 1
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "environment",
            "ticketURI"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}`,
	"https://cdevents.dev/0.3.0/schema/incident-resolved-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/incident-resolved-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.incident.resolved.0.1.0"
          ],
          "default": "dev.cdevents.incident.resolved.0.1.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "incident"
          ],
          "default": "incident"
        },
        "content": {
          "properties": {
            "description": {
              "type": "string"
            },
            "environment": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "service": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "artifactId": {
              "type": "string",
              "minLength": 1
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "environment"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}`,
	"https://cdevents.dev/0.3.0/schema/pipeline-run-finished-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/pipeline-run-finished-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.pipelinerun.finished.0.1.1"
          ],
          "default": "dev.cdevents.pipelinerun.finished.0.1.1"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "pipelineRun"
          ],
          "default": "pipelineRun"
        },
        "content": {
          "properties": {
            "pipelineName": {
              "type": "string"
            },
            "url": {
              "type": "string"
            },
            "outcome": {
              "type": "string"
            },
            "errors": {
              "type": "string"
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}`,
	"https://cdevents.dev/0.3.0/schema/pipeline-run-queued-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/pipeline-run-queued-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.pipelinerun.queued.0.1.1"
          ],
          "default": "dev.cdevents.pipelinerun.queued.0.1.1"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "pipelineRun"
          ],
          "default": "pipelineRun"
        },
        "content": {
          "properties": {
            "pipelineName": {
              "type": "string"
            },
            "url": {
              "type": "string"
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}`,
	"https://cdevents.dev/0.3.0/schema/pipeline-run-started-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/pipeline-run-started-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.pipelinerun.started.0.1.1"
          ],
          "default": "dev.cdevents.pipelinerun.started.0.1.1"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "pipelineRun"
          ],
          "default": "pipelineRun"
        },
        "content": {
          "properties": {
            "pipelineName": {
              "type": "string"
            },
            "url": {
              "type": "string"
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "pipelineName",
            "url"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}`,
	"https://cdevents.dev/0.3.0/schema/repository-created-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/repository-created-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.repository.created.0.1.1"
          ],
          "default": "dev.cdevents.repository.created.0.1.1"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "repository"
          ],
          "default": "repository"
        },
        "content": {
          "properties": {
            "name": {
              "type": "string",
              "minLength": 1
            },
            "owner": {
              "type": "string"
            },
            "url": {
              "type": "string",
              "minLength": 1
            },
            "viewUrl": {
              "type": "string"
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "name",
            "url"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}`,
	"https://cdevents.dev/0.3.0/schema/repository-deleted-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/repository-deleted-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.repository.deleted.0.1.1"
          ],
          "default": "dev.cdevents.repository.deleted.0.1.1"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "repository"
          ],
          "default": "repository"
        },
        "content": {
          "properties": {
            "name": {
              "type": "string"
            },
            "owner": {
              "type": "string"
            },
            "url": {
              "type": "string"
            },
            "viewUrl": {
              "type": "string"
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}`,
	"https://cdevents.dev/0.3.0/schema/repository-modified-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/repository-modified-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.repository.modified.0.1.1"
          ],
          "default": "dev.cdevents.repository.modified.0.1.1"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "repository"
          ],
          "default": "repository"
        },
        "content": {
          "properties": {
            "name": {
              "type": "string"
            },
            "owner": {
              "type": "string"
            },
            "url": {
              "type": "string"
            },
            "viewUrl": {
              "type": "string"
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}`,
	"https://cdevents.dev/0.3.0/schema/service-deployed-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/service-deployed-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.service.deployed.0.1.1"
          ],
          "default": "dev.cdevents.service.deployed.0.1.1"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "service"
          ],
          "default": "service"
        },
        "content": {
          "properties": {
            "environment": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "artifactId": {
              "type": "string",
              "minLength": 1
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "environment",
            "artifactId"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
      "type": "string",
      "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}`,
	"https://cdevents.dev/0.3.0/schema/service-published-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/service-published-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.service.published.0.1.1"
          ],
          "default": "dev.cdevents.service.published.0.1.1"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "service"
          ],
          "default": "service"
        },
        "content": {
          "properties": {
            "environment": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}`,
	"https://cdevents.dev/0.3.0/schema/service-removed-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/service-removed-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.service.removed.0.1.1"
          ],
          "default": "dev.cdevents.service.removed.0.1.1"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "service"
          ],
          "default": "service"
        },
        "content": {
          "properties": {
            "environment": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}`,
	"https://cdevents.dev/0.3.0/schema/service-rolledback-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/service-rolledback-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.service.rolledback.0.1.1"
          ],
          "default": "dev.cdevents.service.rolledback.0.1.1"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "service"
          ],
          "default": "service"
        },
        "content": {
          "properties": {
            "environment": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "artifactId": {
              "type": "string",
              "minLength": 1
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "environment",
            "artifactId"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}`,
	"https://cdevents.dev/0.3.0/schema/service-upgraded-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/service-upgraded-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.service.upgraded.0.1.1"
          ],
          "default": "dev.cdevents.service.upgraded.0.1.1"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "service"
          ],
          "default": "service"
        },
        "content": {
          "properties": {
            "environment": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "artifactId": {
              "type": "string",
              "minLength": 1
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "environment",
            "artifactId"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}`,
	"https://cdevents.dev/0.3.0/schema/task-run-finished-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/task-run-finished-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.taskrun.finished.0.1.1"
          ],
          "default": "dev.cdevents.taskrun.finished.0.1.1"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "taskRun"
          ],
          "default": "taskRun"
        },
        "content": {
          "properties": {
            "taskName": {
              "type": "string"
            },
            "url": {
              "type": "string"
            },
            "pipelineRun": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "outcome": {
              "type": "string"
            },
            "errors": {
              "type": "string"
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}`,
	"https://cdevents.dev/0.3.0/schema/task-run-started-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/task-run-started-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.taskrun.started.0.1.1"
          ],
          "default": "dev.cdevents.taskrun.started.0.1.1"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "taskRun"
          ],
          "default": "taskRun"
        },
        "content": {
          "properties": {
            "taskName": {
              "type": "string"
            },
            "url": {
              "type": "string"
            },
            "pipelineRun": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}`,
	"https://cdevents.dev/0.3.0/schema/test-case-run-finished-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/test-case-run-finished-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.testcaserun.finished.0.1.0"
          ],
          "default": "dev.cdevents.testcaserun.finished.0.1.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "testCaseRun"
          ],
          "default": "testCaseRun"
        },
        "content": {
          "properties": {
            "outcome": {
              "type": "string",
              "enum": [
                "pass",
                "fail",
                "cancel",
                "error"
              ]
            },
            "severity": {
              "type": "string",
              "enum": [
                "low",
                "medium",
                "high",
                "critical"
              ]
            },
            "reason": {
              "type": "string"
            },
            "environment": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "testSuiteRun": {
              "type": "object",
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string"
                }
              },
              "additionalProperties": false,
              "required": [
                "id"
              ]
            },
            "testCase": {
              "type": "object",
              "additionalProperties": false,
              "required": [
                "id"
              ],
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "version": {
                  "type": "string"
                },
                "name": {
                  "type": "string"
                },
                "type": {
                  "type": "string",
                  "enum": [
                    "performance",
                    "functional",
                    "unit",
                    "security",
                    "compliance",
                    "integration",
                    "e2e",
                    "other"
                  ]
                },
                "uri": {
                  "type": "string",
                  "format": "uri"
                }
              }
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "outcome",
            "environment"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.3.0/schema/test-case-run-queued-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/test-case-run-queued-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.testcaserun.queued.0.1.0"
          ],
          "default": "dev.cdevents.testcaserun.queued.0.1.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "testCaseRun"
          ],
          "default": "testCaseRun"
        },
        "content": {
          "properties": {
            "trigger": {
              "type": "object",
              "properties": {
                "type": {
                  "type": "string",
                  "enum": [
                    "manual",
                    "pipeline",
                    "event",
                    "schedule",
                    "other"
                  ]
                },
                "uri": {
                  "type": "string",
                  "format": "uri"
                }
              }
            },
            "environment": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "testSuiteRun": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "testCase": {
              "type": "object",
              "additionalProperties": false,
              "required": [
                "id"
              ],
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "version": {
                  "type": "string"
                },
                "name": {
                  "type": "string"
                },
                "type": {
                  "type": "string",
                  "enum": [
                    "performance",
                    "functional",
                    "unit",
                    "security",
                    "compliance",
                    "integration",
                    "e2e",
                    "other"
                  ]
                },
                "uri": {
                  "type": "string",
                  "format": "uri"
                }
              }
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "environment"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.3.0/schema/test-case-run-started-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/test-case-run-started-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.testcaserun.started.0.1.0"
          ],
          "default": "dev.cdevents.testcaserun.started.0.1.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "testCaseRun"
          ],
          "default": "testCaseRun"
        },
        "content": {
          "properties": {
            "trigger": {
              "type": "object",
              "properties": {
                "type": {
                  "type": "string",
                  "enum": [
                    "manual",
                    "pipeline",
                    "event",
                    "schedule",
                    "other"
                  ]
                },
                "uri": {
                  "type": "string",
                  "format": "uri"
                }
              }
            },
            "environment": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "testSuiteRun": {
              "type": "object",
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string"
                }
              },
              "additionalProperties": false,
              "required": [
                "id"
              ]
            },
            "testCase": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "version": {
                  "type": "string"
                },
                "name": {
                  "type": "string"
                },
                "type": {
                  "type": "string",
                  "enum": [
                    "performance",
                    "functional",
                    "unit",
                    "security",
                    "compliance",
                    "integration",
                    "e2e",
                    "other"
                  ]
                },
                "uri": {
                  "type": "string",
                  "format": "uri"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "environment"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.3.0/schema/test-output-published-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/test-output-published-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.testoutput.published.0.1.0"
          ],
          "default": "dev.cdevents.testoutput.published.0.1.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "testOutput"
          ],
          "default": "testOutput"
        },
        "content": {
          "properties": {
            "outputType": {
              "type": "string",
              "enum": [
                "report",
                "video",
                "image",
                "log",
                "other"
              ]
            },
            "format": {
              "type": "string",
              "example": "application/pdf"
            },
            "uri": {
              "type": "string",
              "format": "uri"
            },
            "testCaseRun": {
              "type": "object",
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string"
                }
              },
              "additionalProperties": false,
              "required": [
                "id"
              ]
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "outputType",
            "format"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.3.0/schema/test-suite-finished-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/test-suite-finished-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.testsuiterun.finished.0.1.0"
          ],
          "default": "dev.cdevents.testsuiterun.finished.0.1.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "testSuiteRun"
          ],
          "default": "testSuiteRun"
        },
        "content": {
          "properties": {
            "environment": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "testSuite": {
              "type": "object",
              "additionalProperties": false,
              "required": [
                "id"
              ],
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "version": {
                  "type": "string"
                },
                "name": {
                  "type": "string"
                },
                "uri": {
                  "type": "string",
                  "format": "uri"
                }
              }
            },
            "outcome": {
              "type": "string",
              "enum": [
                "pass",
                "fail",
                "cancel",
                "error"
              ]
            },
            "severity": {
              "type": "string",
              "enum": [
                "low",
                "medium",
                "high",
                "critical"
              ]
            },
            "reason": {
              "type": "string"
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "outcome",
            "environment"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.3.0/schema/test-suite-run-queued-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/test-suite-run-queued-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.testsuiterun.queued.0.1.0"
          ],
          "default": "dev.cdevents.testsuiterun.queued.0.1.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "testSuiteRun"
          ],
          "default": "testSuiteRun"
        },
        "content": {
          "properties": {
            "trigger": {
              "type": "object",
              "properties": {
                "type": {
                  "type": "string",
                  "enum": [
                    "manual",
                    "pipeline",
                    "event",
                    "schedule",
                    "other"
                  ]
                },
                "uri": {
                  "type": "string",
                  "format": "uri"
                }
              }
            },
            "environment": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "testSuite": {
              "type": "object",
              "additionalProperties": false,
              "required": [
                "id"
              ],
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "version": {
                  "type": "string"
                },
                "name": {
                  "type": "string"
                },
                "url": {
                  "type": "string",
                  "format": "uri"
                }
              }
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "environment"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.3.0/schema/test-suite-run-started-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.3.0/schema/test-suite-run-started-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.testsuiterun.started.0.1.0"
          ],
          "default": "dev.cdevents.testsuiterun.started.0.1.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "testSuiteRun"
          ],
          "default": "testSuiteRun"
        },
        "content": {
          "properties": {
            "trigger": {
              "type": "object",
              "properties": {
                "type": {
                  "type": "string",
                  "enum": [
                    "manual",
                    "pipeline",
                    "event",
                    "schedule",
                    "other"
                  ]
                },
                "uri": {
                  "type": "string",
                  "format": "uri"
                }
              }
            },
            "environment": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "testSuite": {
              "type": "object",
              "additionalProperties": false,
              "required": [
                "id"
              ],
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "version": {
                  "type": "string"
                },
                "name": {
                  "type": "string"
                },
                "uri": {
                  "type": "string",
                  "format": "uri"
                }
              }
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "environment"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/artifact-deleted-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/artifact-deleted-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.artifact.deleted.0.1.0"
          ],
          "default": "dev.cdevents.artifact.deleted.0.1.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "artifact"
          ],
          "default": "artifact"
        },
        "content": {
          "properties": {
            "user": {
              "type": "string",
              "minLength": 1
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/artifact-downloaded-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/artifact-downloaded-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.artifact.downloaded.0.1.0"
          ],
          "default": "dev.cdevents.artifact.downloaded.0.1.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "artifact"
          ],
          "default": "artifact"
        },
        "content": {
          "properties": {
            "user": {
              "type": "string",
              "minLength": 1
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/artifact-packaged-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/artifact-packaged-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.artifact.packaged.0.2.0"
          ],
          "default": "dev.cdevents.artifact.packaged.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "artifact"
          ],
          "default": "artifact"
        },
        "content": {
          "properties": {
            "change": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "sbom": {
              "properties": {
                "uri": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "uri"
              ]
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "change"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/artifact-published-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/artifact-published-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.artifact.published.0.2.0"
          ],
          "default": "dev.cdevents.artifact.published.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "artifact"
          ],
          "default": "artifact"
        },
        "content": {
          "properties": {
            "sbom": {
              "properties": {
                "uri": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "uri"
              ]
            },
            "user": {
              "type": "string",
              "minLength": 1
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/artifact-signed-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/artifact-signed-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.artifact.signed.0.2.0"
          ],
          "default": "dev.cdevents.artifact.signed.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "artifact"
          ],
          "default": "artifact"
        },
        "content": {
          "properties": {
            "signature": {
              "type": "string",
              "minLength": 1
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "signature"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/branch-created-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/branch-created-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.branch.created.0.2.0"
          ],
          "default": "dev.cdevents.branch.created.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "branch"
          ],
          "default": "branch"
        },
        "content": {
          "properties": {
            "repository": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/branch-deleted-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/branch-deleted-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.branch.deleted.0.2.0"
          ],
          "default": "dev.cdevents.branch.deleted.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "branch"
          ],
          "default": "branch"
        },
        "content": {
          "properties": {
            "repository": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/build-finished-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/build-finished-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.build.finished.0.2.0"
          ],
          "default": "dev.cdevents.build.finished.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "build"
          ],
          "default": "build"
        },
        "content": {
          "properties": {
            "artifactId": {
              "type": "string"
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/build-queued-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/build-queued-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.build.queued.0.2.0"
          ],
          "default": "dev.cdevents.build.queued.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "build"
          ],
          "default": "build"
        },
        "content": {
          "properties": {},
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/build-started-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/build-started-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.build.started.0.2.0"
          ],
          "default": "dev.cdevents.build.started.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "build"
          ],
          "default": "build"
        },
        "content": {
          "properties": {},
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/change-abandoned-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/change-abandoned-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.change.abandoned.0.2.0"
          ],
          "default": "dev.cdevents.change.abandoned.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "change"
          ],
          "default": "change"
        },
        "content": {
          "properties": {
            "repository": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/change-created-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/change-created-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.change.created.0.3.0"
          ],
          "default": "dev.cdevents.change.created.0.3.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "change"
          ],
          "default": "change"
        },
        "content": {
          "properties": {
            "description": {
              "type": "string",
              "minLength": 1
            },
            "repository": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/change-merged-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/change-merged-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.change.merged.0.2.0"
          ],
          "default": "dev.cdevents.change.merged.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "change"
          ],
          "default": "change"
        },
        "content": {
          "properties": {
            "repository": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/change-reviewed-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/change-reviewed-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.change.reviewed.0.2.0"
          ],
          "default": "dev.cdevents.change.reviewed.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "change"
          ],
          "default": "change"
        },
        "content": {
          "properties": {
            "repository": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/change-updated-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/change-updated-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.change.updated.0.2.0"
          ],
          "default": "dev.cdevents.change.updated.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "change"
          ],
          "default": "change"
        },
        "content": {
          "properties": {
            "repository": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/custom": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/custom",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "pattern": "^dev\\.cdeventsx\\.[a-zA-Z0-9]+-[a-zA-Z]+\\.[a-zA-Z]+\\.[0-9]\\.[0-9]\\.[0-9]$"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "pattern": "^[a-zA-Z0-9]+-[a-zA-Z]+$"
        },
        "content": {
          "type": "object",
          "additionalProperties": true
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}`,
	"https://cdevents.dev/0.4.1/schema/environment-created-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/environment-created-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.environment.created.0.2.0"
          ],
          "default": "dev.cdevents.environment.created.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "environment"
          ],
          "default": "environment"
        },
        "content": {
          "properties": {
            "name": {
              "type": "string"
            },
            "url": {
              "type": "string"
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/environment-deleted-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/environment-deleted-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.environment.deleted.0.2.0"
          ],
          "default": "dev.cdevents.environment.deleted.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "environment"
          ],
          "default": "environment"
        },
        "content": {
          "properties": {
            "name": {
              "type": "string"
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/environment-modified-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/environment-modified-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.environment.modified.0.2.0"
          ],
          "default": "dev.cdevents.environment.modified.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "environment"
          ],
          "default": "environment"
        },
        "content": {
          "properties": {
            "name": {
              "type": "string"
            },
            "url": {
              "type": "string"
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/incident-detected-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/incident-detected-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.incident.detected.0.2.0"
          ],
          "default": "dev.cdevents.incident.detected.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "incident"
          ],
          "default": "incident"
        },
        "content": {
          "properties": {
            "description": {
              "type": "string"
            },
            "environment": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "service": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "artifactId": {
              "type": "string",
              "minLength": 1
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "environment"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/incident-reported-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/incident-reported-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.incident.reported.0.2.0"
          ],
          "default": "dev.cdevents.incident.reported.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "incident"
          ],
          "default": "incident"
        },
        "content": {
          "properties": {
            "description": {
              "type": "string"
            },
            "environment": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "ticketURI": {
              "type": "string",
              "format": "uri",
              "minLength": 1
            },
            "service": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "artifactId": {
              "type": "string",
              "minLength": 1
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "environment",
            "ticketURI"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/incident-resolved-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/incident-resolved-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.incident.resolved.0.2.0"
          ],
          "default": "dev.cdevents.incident.resolved.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "incident"
          ],
          "default": "incident"
        },
        "content": {
          "properties": {
            "description": {
              "type": "string"
            },
            "environment": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "service": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "artifactId": {
              "type": "string",
              "minLength": 1
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "environment"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/links/embeddedlinkend": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/links/embeddedlinkend",
  "properties": {
    "linkType": {
      "type": "string",
      "enum": [
        "END"
      ]
    },
    "from": {
      "description": "When consuming a CDEvent, you are consuming a parent event. So, when looking at the 'from' key, this is the parent's parent.",
      "type": "object",
      "properties": {
        "contextId": {
          "type": "string",
          "minLength": 1
        }
      },
      "required": [
        "contextId"
      ]
    },
    "tags": {
      "type": "object",
      "additionalProperties": true
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "linkType"
  ]
}

`,
	"https://cdevents.dev/0.4.1/schema/links/embeddedlinkpath": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/links/embeddedlinkpath",
  "properties": {
    "linkType": {
      "type": "string",
      "enum": [
        "PATH"
      ]
    },
    "from": {
      "description": "When consuming a CDEvent, you are consuming a parent event. So, when looking at the 'from' key, this is the parent's parent.",
      "type": "object",
      "properties": {
        "contextId": {
          "type": "string",
          "minLength": 1
        }
      },
      "required": [
        "contextId"
      ]
    },
    "tags": {
      "type": "object",
      "additionalProperties": true
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "linkType",
    "from"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/links/embeddedlinkrelation": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/links/embeddedlinkrelation",
  "properties": {
    "linkType": {
      "type": "string",
      "enum": [
        "RELATION"
      ]
    },
    "linkKind": {
      "type": "string",
      "minLength": 1
    },
    "target": {
      "description": "",
      "type": "object",
      "properties": {
        "contextId": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "tags": {
      "type": "object",
      "additionalProperties": true
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "linkType",
    "linkKind",
    "target"
  ]
}

`,
	"https://cdevents.dev/0.4.1/schema/links/embeddedlinksarray": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/links/embeddedlinksarray",
  "type": "array",
  "items": {
    "anyOf": [
      {
        "type": "object",
        "$ref": "embeddedlinkend"
      },
      {
        "type": "object",
        "$ref": "embeddedlinkpath"
      },
      {
        "type": "object",
        "$ref": "embeddedlinkrelation"
      }
    ]
  }
}
`,
	"https://cdevents.dev/0.4.1/schema/links/linkend": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/links/linkend",
  "properties": {
    "chainId": {
      "description": "This represents the full lifecycles of a series of events in CDEvents",
      "type": "string",
      "minLength": 1
    },
    "linkType": {
      "description": "The type associated with the link. In this case, 'END', suggesting the end of some CI/CD lifecycle",
      "type": "string",
      "enum": [
        "END"
      ]
    },
    "timestamp": {
      "type": "string",
      "format": "date-time"
    },
    "from": {
      "description": "This is the context ID of the producing CDEvent.",
      "type": "object",
      "properties": {
        "contextId": {
          "type": "string",
          "minLength": 1
        }
      },
      "required": [
        "contextId"
      ]
    },
    "end": {
      "description": "This is the context ID of the final CDEvent in the chain",
      "type": "object",
      "properties": {
        "contextId": {
          "type": "string",
          "minLength": 1
        }
      },
      "required": [
        "contextId"
      ]
    },
    "tags": {
      "type": "object",
      "additionalProperties": true
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "chainId",
    "linkType",
    "timestamp",
    "from",
    "end"
  ]
}

`,
	"https://cdevents.dev/0.4.1/schema/links/linkpath": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/links/linkpath",
  "properties": {
    "chainId": {
      "type": "string",
      "minLength": 1
    },
    "timestamp": {
      "type": "string",
      "format": "date-time"
    },
    "linkType": {
      "type": "string",
      "enum": [
        "PATH"
      ]
    },
    "from": {
      "type": "object",
      "properties": {
        "contextId": {
          "type": "string",
          "minLength": 1
        }
      },
      "required": [
        "contextId"
      ]
    },
    "to": {
      "type": "object",
      "properties": {
        "contextId": {
          "type": "string",
          "minLength": 1
        }
      },
      "required": [
        "contextId"
      ]
    },
    "tags": {
      "type": "object",
      "additionalProperties": true
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "chainId",
    "linkType",
    "timestamp",
    "from",
    "to"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/links/linkrelation": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/links/linkrelation",
  "properties": {
    "chainId": {
      "description": "This represents the full lifecycles of a series of events in CDEvents",
      "type": "string",
      "minLength": 1
    },
    "linkType": {
      "type": "string",
      "enum": [
        "RELATION"
      ]
    },
    "linkKind": {
      "type": "string",
      "minLength": 1
    },
    "timestamp": {
      "type": "string",
      "format": "date-time"
    },
    "source": {
      "description": "",
      "type": "object",
      "properties": {
        "contextId": {
          "type": "string",
          "minLength": 1
        }
      },
      "required": [
        "contextId"
      ]
    },
    "target": {
      "description": "",
      "type": "object",
      "properties": {
        "contextId": {
          "type": "string",
          "minLength": 1
        }
      },
      "required": [
        "contextId"
      ]
    },
    "tags": {
      "type": "object",
      "additionalProperties": true
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "chainId",
    "linkType",
    "timestamp",
    "source",
    "target"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/links/linkstart": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/links/linkstart",
  "properties": {
    "chainId": {
      "description": "This represents the full lifecycles of a series of events in CDEvents",
      "type": "string",
      "minLength": 1
    },
    "linkType": {
      "description": "The type associated with the link. In this case, 'START', suggesting the start of some CI/CD lifecycle",
      "type": "string",
      "enum": [
        "START"
      ]
    },
    "timestamp": {
      "type": "string",
      "format": "date-time"
    },
    "start": {
      "description": "This is the context ID of the starting CDEvent in the chain.",
      "type": "object",
      "properties": {
        "contextId": {
          "type": "string",
          "minLength": 1
        }
      },
      "required": [
        "contextId"
      ]
    },
    "tags": {
      "type": "object",
      "additionalProperties": true
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "chainId",
    "linkType",
    "timestamp",
    "start"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/pipeline-run-finished-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/pipeline-run-finished-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.pipelinerun.finished.0.2.0"
          ],
          "default": "dev.cdevents.pipelinerun.finished.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "pipelineRun"
          ],
          "default": "pipelineRun"
        },
        "content": {
          "properties": {
            "pipelineName": {
              "type": "string"
            },
            "url": {
              "type": "string"
            },
            "outcome": {
              "type": "string"
            },
            "errors": {
              "type": "string"
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/pipeline-run-queued-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/pipeline-run-queued-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.pipelinerun.queued.0.2.0"
          ],
          "default": "dev.cdevents.pipelinerun.queued.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "pipelineRun"
          ],
          "default": "pipelineRun"
        },
        "content": {
          "properties": {
            "pipelineName": {
              "type": "string"
            },
            "url": {
              "type": "string"
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/pipeline-run-started-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/pipeline-run-started-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.pipelinerun.started.0.2.0"
          ],
          "default": "dev.cdevents.pipelinerun.started.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "pipelineRun"
          ],
          "default": "pipelineRun"
        },
        "content": {
          "properties": {
            "pipelineName": {
              "type": "string"
            },
            "url": {
              "type": "string"
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "pipelineName",
            "url"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/repository-created-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/repository-created-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.repository.created.0.2.0"
          ],
          "default": "dev.cdevents.repository.created.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "repository"
          ],
          "default": "repository"
        },
        "content": {
          "properties": {
            "name": {
              "type": "string",
              "minLength": 1
            },
            "owner": {
              "type": "string"
            },
            "url": {
              "type": "string",
              "minLength": 1
            },
            "viewUrl": {
              "type": "string"
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "name",
            "url"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/repository-deleted-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/repository-deleted-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.repository.deleted.0.2.0"
          ],
          "default": "dev.cdevents.repository.deleted.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "repository"
          ],
          "default": "repository"
        },
        "content": {
          "properties": {
            "name": {
              "type": "string"
            },
            "owner": {
              "type": "string"
            },
            "url": {
              "type": "string"
            },
            "viewUrl": {
              "type": "string"
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/repository-modified-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/repository-modified-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.repository.modified.0.2.0"
          ],
          "default": "dev.cdevents.repository.modified.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "repository"
          ],
          "default": "repository"
        },
        "content": {
          "properties": {
            "name": {
              "type": "string"
            },
            "owner": {
              "type": "string"
            },
            "url": {
              "type": "string"
            },
            "viewUrl": {
              "type": "string"
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/service-deployed-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/service-deployed-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.service.deployed.0.2.0"
          ],
          "default": "dev.cdevents.service.deployed.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "service"
          ],
          "default": "service"
        },
        "content": {
          "properties": {
            "environment": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "artifactId": {
              "type": "string",
              "minLength": 1
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "environment",
            "artifactId"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/service-published-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/service-published-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.service.published.0.2.0"
          ],
          "default": "dev.cdevents.service.published.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "service"
          ],
          "default": "service"
        },
        "content": {
          "properties": {
            "environment": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/service-removed-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/service-removed-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.service.removed.0.2.0"
          ],
          "default": "dev.cdevents.service.removed.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "service"
          ],
          "default": "service"
        },
        "content": {
          "properties": {
            "environment": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/service-rolledback-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/service-rolledback-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.service.rolledback.0.2.0"
          ],
          "default": "dev.cdevents.service.rolledback.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "service"
          ],
          "default": "service"
        },
        "content": {
          "properties": {
            "environment": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "artifactId": {
              "type": "string",
              "minLength": 1
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "environment",
            "artifactId"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/service-upgraded-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/service-upgraded-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.service.upgraded.0.2.0"
          ],
          "default": "dev.cdevents.service.upgraded.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "service"
          ],
          "default": "service"
        },
        "content": {
          "properties": {
            "environment": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "artifactId": {
              "type": "string",
              "minLength": 1
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "environment",
            "artifactId"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/task-run-finished-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/task-run-finished-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.taskrun.finished.0.2.0"
          ],
          "default": "dev.cdevents.taskrun.finished.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "taskRun"
          ],
          "default": "taskRun"
        },
        "content": {
          "properties": {
            "taskName": {
              "type": "string"
            },
            "url": {
              "type": "string"
            },
            "pipelineRun": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "outcome": {
              "type": "string"
            },
            "errors": {
              "type": "string"
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/task-run-started-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/task-run-started-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.taskrun.started.0.2.0"
          ],
          "default": "dev.cdevents.taskrun.started.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "taskRun"
          ],
          "default": "taskRun"
        },
        "content": {
          "properties": {
            "taskName": {
              "type": "string"
            },
            "url": {
              "type": "string"
            },
            "pipelineRun": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            }
          },
          "additionalProperties": false,
          "type": "object"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/test-case-run-finished-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/test-case-run-finished-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.testcaserun.finished.0.2.0"
          ],
          "default": "dev.cdevents.testcaserun.finished.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "testCaseRun"
          ],
          "default": "testCaseRun"
        },
        "content": {
          "properties": {
            "outcome": {
              "type": "string",
              "enum": [
                "pass",
                "fail",
                "cancel",
                "error"
              ]
            },
            "severity": {
              "type": "string",
              "enum": [
                "low",
                "medium",
                "high",
                "critical"
              ]
            },
            "reason": {
              "type": "string"
            },
            "environment": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "testSuiteRun": {
              "type": "object",
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string"
                }
              },
              "additionalProperties": false,
              "required": [
                "id"
              ]
            },
            "testCase": {
              "type": "object",
              "additionalProperties": false,
              "required": [
                "id"
              ],
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "version": {
                  "type": "string"
                },
                "name": {
                  "type": "string"
                },
                "type": {
                  "type": "string",
                  "enum": [
                    "performance",
                    "functional",
                    "unit",
                    "security",
                    "compliance",
                    "integration",
                    "e2e",
                    "other"
                  ]
                },
                "uri": {
                  "type": "string",
                  "format": "uri"
                }
              }
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "outcome",
            "environment"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/test-case-run-queued-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/test-case-run-queued-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.testcaserun.queued.0.2.0"
          ],
          "default": "dev.cdevents.testcaserun.queued.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "testCaseRun"
          ],
          "default": "testCaseRun"
        },
        "content": {
          "properties": {
            "trigger": {
              "type": "object",
              "properties": {
                "type": {
                  "type": "string",
                  "enum": [
                    "manual",
                    "pipeline",
                    "event",
                    "schedule",
                    "other"
                  ]
                },
                "uri": {
                  "type": "string",
                  "format": "uri"
                }
              }
            },
            "environment": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "testSuiteRun": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "testCase": {
              "type": "object",
              "additionalProperties": false,
              "required": [
                "id"
              ],
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "version": {
                  "type": "string"
                },
                "name": {
                  "type": "string"
                },
                "type": {
                  "type": "string",
                  "enum": [
                    "performance",
                    "functional",
                    "unit",
                    "security",
                    "compliance",
                    "integration",
                    "e2e",
                    "other"
                  ]
                },
                "uri": {
                  "type": "string",
                  "format": "uri"
                }
              }
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "environment"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/test-case-run-skipped-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/test-case-run-skipped-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.testcaserun.skipped.0.1.0"
          ],
          "default": "dev.cdevents.testcaserun.skipped.0.1.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "testCaseRun"
          ],
          "default": "testCaseRun"
        },
        "content": {
          "properties": {
            "reason": {
              "type": "string"
            },
            "environment": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "testSuiteRun": {
              "type": "object",
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string"
                }
              },
              "additionalProperties": false,
              "required": [
                "id"
              ]
            },
            "testCase": {
              "type": "object",
              "additionalProperties": false,
              "required": [
                "id"
              ],
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "version": {
                  "type": "string"
                },
                "name": {
                  "type": "string"
                },
                "type": {
                  "type": "string",
                  "enum": [
                    "performance",
                    "functional",
                    "unit",
                    "security",
                    "compliance",
                    "integration",
                    "e2e",
                    "other"
                  ]
                },
                "uri": {
                  "type": "string",
                  "format": "uri"
                }
              }
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": []
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}`,
	"https://cdevents.dev/0.4.1/schema/test-case-run-started-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/test-case-run-started-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.testcaserun.started.0.2.0"
          ],
          "default": "dev.cdevents.testcaserun.started.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "testCaseRun"
          ],
          "default": "testCaseRun"
        },
        "content": {
          "properties": {
            "trigger": {
              "type": "object",
              "properties": {
                "type": {
                  "type": "string",
                  "enum": [
                    "manual",
                    "pipeline",
                    "event",
                    "schedule",
                    "other"
                  ]
                },
                "uri": {
                  "type": "string",
                  "format": "uri"
                }
              }
            },
            "environment": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "testSuiteRun": {
              "type": "object",
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string"
                }
              },
              "additionalProperties": false,
              "required": [
                "id"
              ]
            },
            "testCase": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "version": {
                  "type": "string"
                },
                "name": {
                  "type": "string"
                },
                "type": {
                  "type": "string",
                  "enum": [
                    "performance",
                    "functional",
                    "unit",
                    "security",
                    "compliance",
                    "integration",
                    "e2e",
                    "other"
                  ]
                },
                "uri": {
                  "type": "string",
                  "format": "uri"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "environment"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/test-output-published-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/test-output-published-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.testoutput.published.0.2.0"
          ],
          "default": "dev.cdevents.testoutput.published.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "testOutput"
          ],
          "default": "testOutput"
        },
        "content": {
          "properties": {
            "outputType": {
              "type": "string",
              "enum": [
                "report",
                "video",
                "image",
                "log",
                "other"
              ]
            },
            "format": {
              "type": "string",
              "example": "application/pdf"
            },
            "uri": {
              "type": "string",
              "format": "uri"
            },
            "testCaseRun": {
              "type": "object",
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string"
                }
              },
              "additionalProperties": false,
              "required": [
                "id"
              ]
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "outputType",
            "format"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/test-suite-run-finished-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/test-suite-run-finished-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.testsuiterun.finished.0.2.0"
          ],
          "default": "dev.cdevents.testsuiterun.finished.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "testSuiteRun"
          ],
          "default": "testSuiteRun"
        },
        "content": {
          "properties": {
            "environment": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "testSuite": {
              "type": "object",
              "additionalProperties": false,
              "required": [
                "id"
              ],
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "version": {
                  "type": "string"
                },
                "name": {
                  "type": "string"
                },
                "uri": {
                  "type": "string",
                  "format": "uri"
                }
              }
            },
            "outcome": {
              "type": "string",
              "enum": [
                "pass",
                "fail",
                "cancel",
                "error"
              ]
            },
            "severity": {
              "type": "string",
              "enum": [
                "low",
                "medium",
                "high",
                "critical"
              ]
            },
            "reason": {
              "type": "string"
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "outcome",
            "environment"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/test-suite-run-queued-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/test-suite-run-queued-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.testsuiterun.queued.0.2.0"
          ],
          "default": "dev.cdevents.testsuiterun.queued.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "testSuiteRun"
          ],
          "default": "testSuiteRun"
        },
        "content": {
          "properties": {
            "trigger": {
              "type": "object",
              "properties": {
                "type": {
                  "type": "string",
                  "enum": [
                    "manual",
                    "pipeline",
                    "event",
                    "schedule",
                    "other"
                  ]
                },
                "uri": {
                  "type": "string",
                  "format": "uri"
                }
              }
            },
            "environment": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "testSuite": {
              "type": "object",
              "additionalProperties": false,
              "required": [
                "id"
              ],
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "version": {
                  "type": "string"
                },
                "name": {
                  "type": "string"
                },
                "url": {
                  "type": "string",
                  "format": "uri"
                }
              }
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "environment"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/test-suite-run-started-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/test-suite-run-started-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.testsuiterun.started.0.2.0"
          ],
          "default": "dev.cdevents.testsuiterun.started.0.2.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "testSuiteRun"
          ],
          "default": "testSuiteRun"
        },
        "content": {
          "properties": {
            "trigger": {
              "type": "object",
              "properties": {
                "type": {
                  "type": "string",
                  "enum": [
                    "manual",
                    "pipeline",
                    "event",
                    "schedule",
                    "other"
                  ]
                },
                "uri": {
                  "type": "string",
                  "format": "uri"
                }
              }
            },
            "environment": {
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "source": {
                  "type": "string",
                  "minLength": 1,
                  "format": "uri-reference"
                }
              },
              "additionalProperties": false,
              "type": "object",
              "required": [
                "id"
              ]
            },
            "testSuite": {
              "type": "object",
              "additionalProperties": false,
              "required": [
                "id"
              ],
              "properties": {
                "id": {
                  "type": "string",
                  "minLength": 1
                },
                "version": {
                  "type": "string"
                },
                "name": {
                  "type": "string"
                },
                "uri": {
                  "type": "string",
                  "format": "uri"
                }
              }
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "environment"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/ticket-closed-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/ticket-closed-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.ticket.closed.0.1.0"
          ],
          "default": "dev.cdevents.ticket.closed.0.1.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "ticket"
          ],
          "default": "ticket"
        },
        "content": {
          "properties": {
            "summary": {
              "type": "string"
            },
            "ticketType": {
              "anyOf": [
                {
                  "type": "string",
                  "enum": [
                    "bug",
                    "enhancement",
                    "incident",
                    "task",
                    "question"
                  ]
                },
                {
                  "type": "string"
                }
              ]
            },
            "group": {
              "type": "string"
            },
            "creator": {
              "type": "string",
              "minLength": 1
            },
            "assignees": {
              "type": "array",
              "items": {
                "type": "string"
              }
            },
            "priority": {
              "anyOf": [
                {
                  "type": "string",
                  "enum": [
                    "low",
                    "medium",
                    "high"
                  ]
                },
                {
                  "type": "string"
                }
              ]
            },
            "labels": {
              "type": "array",
              "items": {
                "type": "string"
              }
            },
            "milestone": {
              "type": "string"
            },
            "uri": {
              "type": "string",
              "minLength": 1,
              "format": "uri-reference"
            },
            "resolution": {
              "anyOf": [
                {
                  "type": "string",
                  "enum": [
                    "completed",
                    "withdrawn",
                    "duplicate"
                  ]
                },
                {
                  "type": "string",
                  "minLength": 1
                }
              ]
            },
            "updatedBy": {
              "type": "string"
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "uri",
            "resolution"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/ticket-created-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/ticket-created-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.ticket.created.0.1.0"
          ],
          "default": "dev.cdevents.ticket.created.0.1.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "ticket"
          ],
          "default": "ticket"
        },
        "content": {
          "properties": {
            "summary": {
              "type": "string"
            },
            "ticketType": {
              "anyOf": [
                {
                  "type": "string",
                  "enum": [
                    "bug",
                    "enhancement",
                    "incident",
                    "task",
                    "question"
                  ]
                },
                {
                  "type": "string"
                }
              ]
            },
            "group": {
              "type": "string"
            },
            "creator": {
              "type": "string",
              "minLength": 1
            },
            "assignees": {
              "type": "array",
              "items": {
                "type": "string"
              }
            },
            "priority": {
              "anyOf": [
                {
                  "type": "string",
                  "enum": [
                    "low",
                    "medium",
                    "high"
                  ]
                },
                {
                  "type": "string"
                }
              ]
            },
            "labels": {
              "type": "array",
              "items": {
                "type": "string"
              }
            },
            "milestone": {
              "type": "string"
            },
            "uri": {
              "type": "string",
              "minLength": 1,
              "format": "uri-reference"
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "summary",
            "creator",
            "uri"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
	"https://cdevents.dev/0.4.1/schema/ticket-updated-event": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://cdevents.dev/0.4.1/schema/ticket-updated-event",
  "properties": {
    "context": {
      "properties": {
        "version": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "enum": [
            "dev.cdevents.ticket.updated.0.1.0"
          ],
          "default": "dev.cdevents.ticket.updated.0.1.0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "schemaUri": {
          "type": "string",
          "minLength": 1,
          "format": "uri"
        },
        "chainId": {
          "type": "string",
          "minLength": 1
        },
        "links": {
          "$ref": "links/embeddedlinksarray"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "version",
        "id",
        "source",
        "type",
        "timestamp"
      ]
    },
    "subject": {
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "source": {
          "type": "string",
          "minLength": 1,
          "format": "uri-reference"
        },
        "type": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "ticket"
          ],
          "default": "ticket"
        },
        "content": {
          "properties": {
            "summary": {
              "type": "string"
            },
            "ticketType": {
              "anyOf": [
                {
                  "type": "string",
                  "enum": [
                    "bug",
                    "enhancement",
                    "incident",
                    "task",
                    "question"
                  ]
                },
                {
                  "type": "string"
                }
              ]
            },
            "group": {
              "type": "string"
            },
            "creator": {
              "type": "string",
              "minLength": 1
            },
            "assignees": {
              "type": "array",
              "items": {
                "type": "string"
              }
            },
            "priority": {
              "anyOf": [
                {
                  "type": "string",
                  "enum": [
                    "low",
                    "medium",
                    "high"
                  ]
                },
                {
                  "type": "string"
                }
              ]
            },
            "labels": {
              "type": "array",
              "items": {
                "type": "string"
              }
            },
            "milestone": {
              "type": "string"
            },
            "updatedBy": {
              "type": "string"
            },
            "uri": {
              "type": "string",
              "minLength": 1,
              "format": "uri-reference"
            }
          },
          "additionalProperties": false,
          "type": "object",
          "required": [
            "uri"
          ]
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "id",
        "type",
        "content"
      ]
    },
    "customData": {
      "oneOf": [
        {
          "type": "object"
        },
        {
          "type": "string",
          "contentEncoding": "base64"
        }
      ]
    },
    "customDataContentType": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "required": [
    "context",
    "subject"
  ]
}
`,
}

func GetSchemaBySpecSubjectPredicate(specVersion, subject, predicate string) (string, string, error) {
	id := fmt.Sprintf(CDEventsSchemaURLTemplate, specVersion, subject, predicate)
	if schemaString, found := SchemasById[id]; found {
		return id, schemaString, nil
	}
	return "", "", fmt.Errorf("event %s/%s not found for spec %s in local schema DB", specVersion, subject, predicate)
}
