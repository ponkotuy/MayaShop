declare var Vue: VueStatic;

interface VueObject{
}

interface VueMethods{
    [n: string]: Function;
}

interface VueAttributes{
    [n: string]: string;
}

interface Callback {
    ():void;
}

interface ValueCallback {
    ( value:any ):void;
}

interface VueDirective {
    bind( value:any );
    update( value:any );
    unbind();
}

interface VueOptions{
    data?:VueObject;
    methods?:VueMethods;
    computed?:VueObject;
    paramAttributes?:string[];
    el?:string;
    template?:string;
    replace?:boolean;
    tagName?:string;
    id?:string;
    className?:string;
    attributes?:VueAttributes;
    created?:Callback;
    ready?:Callback;
    attached?:Callback;
    detached?:Callback;
    beforeDestroy?:Callback;
    afterDestroy?:Callback;
    directives?:VueObject;
    filters?:VueObject;
    components?:VueObject;
    partials?:VueObject;
    transitions?:VueObject;
    parent?:VueModel;
    lazy?:boolean;
}

interface VueModel{
    $el:HTMLElement;
    $data:VueObject;
    $options:VueObject;
    $:VueObject;
    $index:VueObject;
    $parent:VueModel;
    $root:VueModel;
    $compiler:VueObject;
    $watch( keypath:string , callback:ValueCallback );
    $unwatch( keypath:string , callback?:ValueCallback );
    $dispatch( event:string , ...args:any[] );
    $broadcast( event:string , ...args:any[] );
    $emit( event:string , ...args:any[] );
    $on( event:string , callback:Callback );
    $once( event:string , callback:Callback );
    $off( event?:string , callback?:Callback );
    $appendTo( selector:string );
    $before( selector:string );
    $after( selector:string );
    $remove();
    $destroy();
}

interface VueStatic{
    new ( options: VueOptions ):VueModel;
    extend( options:VueOptions ):VueStatic;
    config( options: VueOptions );
    config( key:string , value?:any );
    directive( id:string , definition:VueDirective );
    directive( id:string , update:ValueCallback );
    filter( id:string , definition:ValueCallback ):string;
}
