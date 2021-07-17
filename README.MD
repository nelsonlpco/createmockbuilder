![cod cov](https://storage.googleapis.com/imagebadges/createmockbuildercov.svg)

# createmockbuilder

<p>Command line tool to convert json files to typescript/javascript mockBuilder classes.</p>

---

#### basic usage: 

<code>
createmockbuilder -f=./author.json -o=./outputdir
</code>

##### Params
 
 | flag | value          | action |
 |------|----------------|--------|
 | -f   | ./sample.json  | path to input json file |
 | -o   | ./output       | path to output class files, default ./|
 | -t   | js / ts          | template, default  ts |

---

#### Example
input:
author.json

```json
{
	"name": "Nelson Oliveira",
	"active": true,
	"contact": {
		"email": "nelsonlpco@gmail.com"
	}
}
```

output: AuthorBuilder.ts

```typescript
class AuthorBuilder {
	Name: string;
	Active: boolean;
	Contact: any;

	constructor() { 
		Name = "";
		Active = false;
		Contact = {};
	}

	withName(value: string) { 
		this.Name = value;
		return this;
	}

	withActive(value: boolean) { 
		this.Active = valuel;
		return this;
	}

	withContact(value: any) { 
		this.Contact = value;
		return this;
	}

	withDefaultValues() {
		this.Name = "Nelson Oliveira";
		this.Active = true;
		this.Contact = {};
	}

	build() { 
		return { 
			name: this.Name,
			active: this.active,
			contact: this.Contact,
		}
	}
}

class ContactBuilder { 
	Email: string;

	constructor() { 
		this.Email = "";
	}

	withEmail(value: string) { 
		this.Email = value;
		return this;
	}

	withDefaultValues() { 
		this.Email = "nelsonlpco@gmail.com";

		return this;
	}

	build() { 
		return { 
			email: this.Email,
		}
	}
}
```
