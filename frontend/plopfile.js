module.exports = function (plop) {
  // Helpers
  plop.setHelper('kebabCase', text =>
    text
      .replace(/([a-z])([A-Z])/g, '$1-$2')
      .replace(/\s+/g, '-')
      .toLowerCase()
  );

  plop.setHelper('pascalCase', text =>
    text.replace(/(^\w|-\w)/g, clear => clear.replace('-', '').toUpperCase())
  );

  plop.setHelper('plural', text => {
    // primitive pluralizer, can be improved
    return text.endsWith('s') ? text : text + 's';
  });

  // Generator
  plop.setGenerator('entity', {
    description: 'Generate an FSD entity',
    prompts: [
      {
        type: 'input',
        name: 'name',
        message: 'Entity name (singular, e.g. topping):',
      },
    ],
    actions: function () {
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
};
