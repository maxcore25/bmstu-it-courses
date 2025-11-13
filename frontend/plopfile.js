// eslint-disable-next-line @typescript-eslint/no-require-imports
const pluralize = require('pluralize');

module.exports = function (plop) {
  // ---------------- HELPERS ----------------

  // kebab-case
  plop.setHelper('kebabCase', text =>
    text
      .replace(/([a-z])([A-Z])/g, '$1-$2')
      .replace(/\s+/g, '-')
      .toLowerCase()
  );

  // camelCase
  plop.setHelper('camelCase', text =>
    text
      .replace(/-./g, x => x[1].toUpperCase())
      .replace(/^\w/, c => c.toLowerCase())
  );

  // PascalCase
  plop.setHelper('pascalCase', text =>
    text.replace(/(^\w|-\w)/g, clear => clear.replace('-', '').toUpperCase())
  );

  // snake_case
  plop.setHelper('snakeCase', text =>
    text
      .replace(/([a-z])([A-Z])/g, '$1_$2')
      .replace(/\s+/g, '_')
      .toLowerCase()
  );

  // UPPERCASE
  plop.setHelper('upperCase', text => text.toUpperCase());

  // pluralize
  plop.setHelper('plural', text => pluralize(text));

  // ---------------- GENERATOR ----------------
  plop.setGenerator('entity', {
    description: 'Generate an FSD entity',
    prompts: [
      {
        type: 'input',
        name: 'name',
        message: 'Entity name (singular, e.g. topping):',
      },
    ],
    actions: function (data) {
      const base = 'src/entities/{{kebabCase name}}';

      return [
        // api
        {
          type: 'add',
          path: `${base}/api/get-{{kebabCase (plural name)}}.api.ts`,
          templateFile: 'plop-templates/entity/api.hbs',
        },
        // schema
        {
          type: 'add',
          path: `${base}/model/{{kebabCase name}}.schema.ts`,
          templateFile: 'plop-templates/entity/schema.hbs',
        },
        // hook
        {
          type: 'add',
          path: `${base}/model/use-get-{{kebabCase (plural name)}}.ts`,
          templateFile: 'plop-templates/entity/hook.hbs',
        },
        // UI
        {
          type: 'add',
          path: `${base}/ui/{{kebabCase name}}-card.tsx`,
          templateFile: 'plop-templates/entity/ui.hbs',
        },
        // index.ts
        {
          type: 'add',
          path: `${base}/index.ts`,
          templateFile: 'plop-templates/entity/index.hbs',
        },
      ];
    },
  });

  plop.setGenerator('feature', {
    description: 'Generate an FSD feature',
    prompts: [
      {
        type: 'input',
        name: 'name',
        message: 'Feature name (singular, e.g. checkout):',
      },
    ],
    actions: function (data) {
      const base = 'src/features/{{kebabCase name}}';

      return [
        // api
        {
          type: 'add',
          path: `${base}/api/get-{{kebabCase (plural name)}}.api.ts`,
          templateFile: 'plop-templates/entity/api.hbs', // reuse entity API template
        },
        // schema
        {
          type: 'add',
          path: `${base}/model/{{kebabCase name}}.schema.ts`,
          templateFile: 'plop-templates/entity/schema.hbs', // reuse entity schema
        },
        // custom feature button hook
        {
          type: 'add',
          path: `${base}/model/use-{{kebabCase name}}-button.ts`,
          templateFile: 'plop-templates/feature/use-button.hbs',
        },
        // UI button
        {
          type: 'add',
          path: `${base}/ui/{{kebabCase name}}-button.tsx`,
          templateFile: 'plop-templates/feature/button-ui.hbs',
        },
        // index.ts
        {
          type: 'add',
          path: `${base}/index.ts`,
          templateFile: 'plop-templates/feature/index.hbs',
        },
      ];
    },
  });
};
